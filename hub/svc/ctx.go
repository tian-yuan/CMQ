package svc

import (
	"net"
	"sync"
	"time"
	"CMQ/hub/proto/mqtt"
	"github.com/sirupsen/logrus"
	"CMQ/hub/topic"
)

const (
	inactiveState = 0
	activeState   = 1
)

var matcher topic.Matcher
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
		logrus.Infof("Reset packetId: %d", i)
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
