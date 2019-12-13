package svc

import (
	"net"
	"sync"
	"time"

	"github.com/micro/go-micro/util/log"
	"github.com/tian-yuan/CMQ/hub/proto/mqtt"
)

const (
	inactiveState = 0
	activeState   = 1
)

var ClientCtxs []ClientCtx

func GetCtxForFd(fd int64) *ClientCtx {
	i := int(fd)

	if i < len(ClientCtxs) {
		return &ClientCtxs[i]
	}
	return nil
}

var (
	ecNone                = 0
	ecClientClose         = 1
	ecWriteError          = 2
	ecReadError           = 3
	ecServerClose         = 4
	ecSessionTimeout      = 5
	ecInvalidProtocol     = 6
	ecBadClientRequest    = 7
	ecServerInternalError = 8
)

var reasons = map[int]string{
	ecNone:                "None",
	ecClientClose:         "Client Close",
	ecWriteError:          "Write Error",
	ecReadError:           "Read Error",
	ecServerClose:         "Server Close",
	ecSessionTimeout:      "Session Timeout",
	ecInvalidProtocol:     "Invalid Protocol",
	ecBadClientRequest:    "Bad Client Request",
	ecServerInternalError: "Server Internal Error",
}

type ClientCtx struct {
	// Mux/Fd will be never modified
	Mux sync.RWMutex
	Fd  int

	// If StopCh is not nil, the ClientCtx is occupied
	StopCh chan struct{}

	// The following fields should be protected by Mux
	// They may be accessed by several goroutines
	// {{

	Conn       net.Conn
	productKey string
	deviceName string
	guid       uint32
	status     int
	currentId  uint16
	ws         bool
	encoder    mqtt.Encoder

	// }} protected by Mux

	PacketMux      sync.Mutex
	pendingPackets map[uint16]chan int

	decoder      mqtt.Decoder
	lastReceived int64
	nextRefresh  int64

	keepalive    int /* second */
	hasConnected bool
	exitCode     int
	err          error

	sessionKey   string
	sessionValue string
}

var extraSizeHint = 16

// called before reuse ClientCtx
func (ctx *ClientCtx) Reset(ws bool, conn net.Conn) {
	ctx.Mux.Lock()

	ctx.StopCh = make(chan struct{})

	ctx.decoder.Reset(ws)
	ctx.encoder.Reset(ws)

	ctx.Conn = conn

	ctx.lastReceived = time.Now().Unix()
	ctx.nextRefresh = time.Now().Unix()
	ctx.currentId = 0

	ctx.keepalive = 0

	ctx.ws = ws
	ctx.hasConnected = false
	ctx.exitCode = 0
	ctx.err = nil

	ctx.sessionKey = ""
	ctx.sessionValue = ""

	ctx.PacketMux.Lock()
	ctx.pendingPackets = make(map[uint16]chan int)
	ctx.PacketMux.Unlock()

	ctx.Mux.Unlock()
}

func (ctx *ClientCtx) GetPacketId() (id uint16) {
	ctx.Mux.Lock()
	defer ctx.Mux.Unlock()

	ctx.currentId += 1
	if ctx.currentId == 0 {
		ctx.currentId = 1
	}
	id = ctx.currentId
	return
}

func (ctx *ClientCtx) AddPendingPacket(packetId uint16) (c chan int) {
	ctx.PacketMux.Lock()
	defer ctx.PacketMux.Unlock()

	c = make(chan int, 1)
	ctx.pendingPackets[packetId] = c
	return
}

const unackMessageNamespace = ""

func (ctx *ClientCtx) AckPacket(packetId uint16, status int) {
	ctx.PacketMux.Lock()
	defer ctx.PacketMux.Unlock()

	if ctx.pendingPackets == nil {
		return
	}

	if val, ok := ctx.pendingPackets[packetId]; ok {
		val <- status
		delete(ctx.pendingPackets, packetId)
	}
}

func (ctx *ClientCtx) resetPendingPacket() {
	ctx.PacketMux.Lock()
	defer ctx.PacketMux.Unlock()

	for i, f := range ctx.pendingPackets {
		f <- rcNetworkError
		log.Infof("Reset packetId: %d", i)
	}
	ctx.pendingPackets = nil
}

func (ctx *ClientCtx) IsClientMatchedAllLocked(deviceName string, productKey string) bool {
	isActive := ctx.status == activeState
	return isActive
}

func (ctx *ClientCtx) IsClientMatchedMasterLocked(deviceName string, productKey string) bool {
	isActive := ctx.status == activeState
	return isActive
}

func (ctx *ClientCtx) Stop() {
	ctx.Mux.Lock()
	defer ctx.Mux.Unlock()

	if ctx.StopCh != nil {
		select {
		case <-ctx.StopCh:
			return
		default:
			close(ctx.StopCh)
			if ctx.Conn != nil {
				ctx.Conn.SetDeadline(time.Now())
			}
		}
	}
}

var (
	rcRequestOK          = 200
	rcGetMessageError    = 400
	rcClientNotOnline    = 404
	rcEncodeError        = 601
	rcNetworkError       = 602
	rcRequstTimeout      = 603
	rcUnknownServerError = 604
	rcRpcContentError    = 605
)

func (ctx *ClientCtx) PublishMessage(deviceName string, productKey string,
	msg string, topic string, qos int8) int {

	log.Infof("Handle publish message, deviceName %s, productKey: %s", deviceName, productKey)

	h := mqtt.NewPublishPacketTypeHeader(qos, defaultDup, defaultRetain)
	pktId := ctx.GetPacketId()

	var data []byte = []byte(msg)
	p := mqtt.NewPublishPacket(h, topic, pktId, data)

	var isOnline bool
	var ret int = rcRequestOK
	var pf chan int = nil

	ctx.Mux.Lock()

	isOnline = ctx.IsClientMatchedMasterLocked(deviceName, productKey)

	if isOnline {
		if qos == 1 {
			pf = ctx.AddPendingPacket(pktId)
		}

		ret = ctx.publish(p, deviceName, productKey)

		if ret != rcRequestOK {
			log.Errorf("publish to client error, %s, %s", deviceName, productKey)
			if qos == 1 {
				ctx.AckPacket(pktId, ret)
			}
		}
	} else {
		log.Errorf("Client is not online: %s, %s", deviceName, productKey)
		ret = rcClientNotOnline
	}

	ctx.Mux.Unlock()

	if ret != rcRequestOK {
		return ret
	}

	if qos == 1 {
		select {
		case ret := <-pf:
			return ret
		case <-time.After(2 * time.Minute):
			ctx.AckPacket(pktId, rcRequstTimeout)
			return rcRequstTimeout
		}
	}

	return rcRequestOK
}

func (ctx *ClientCtx) publish(p *mqtt.PublishPacket, deviceName string, productKey string) int {
	log.Infof("publish message to client : %s, payload : %s", p.Topic, string(p.Payload))
	if e := p.Encode(&ctx.encoder); e != nil {
		ctx.encoder.ResetState()
		log.Errorf("Encode payload error: %v, %s, %s", e, deviceName, productKey)
		return rcEncodeError
	}

	if e := ctx.encoder.WriteTo(ctx.Conn, 2*time.Minute); e != nil {
		log.Errorf("Push message error: %v, %s, %s", e, deviceName, productKey)
		return rcNetworkError
	}

	return rcRequestOK
}
