package svc

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/util/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/tian-yuan/CMQ/hub/proto/mqtt"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/tian-yuan/iot-common/util"
	"golang.org/x/net/context"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func handleConnection(s *MqttSvc, c net.Conn, ws bool) {
	fd := util.GetConnFd(c)

	log.Infof("receive connection fd : %d", fd)
	ctx := GetCtxForFd(fd)
	if ctx == nil {
		c.Close()
		return
	}

	log.Infof("Client connect, address: %s, isWs: %t, fd: %d", c.RemoteAddr(), ws, fd)

	ctx.Reset(ws, c)

	ctx.handleClient()
	ctx.Mux.Lock()

	log.Infof("Client disconnect, address: %s, isWs: %t, fd: %d, reason: %s, err: %s",
		c.RemoteAddr(), ws, fd, reasons[ctx.exitCode], ctx.err)

	ctx.destroySession()
	ctx.status = inactiveState

	ctx.Conn.Close()
	ctx.Conn = nil
	ctx.StopCh = nil

	ctx.resetPendingPacket()

	ctx.Mux.Unlock()
}

func (ctx *ClientCtx) destroySession() {
	// clear client redis cache info
}

const leastReadTimeout = 120 * time.Second

func (ctx *ClientCtx) handleClient() {
	var err error
	de := &ctx.decoder

	for {
		select {
		case <-ctx.StopCh:
			ctx.exitCode = ecServerClose
			return
		default:
		}

		var timeout time.Duration
		if ctx.keepalive == 0 { // if not receive connect packet
			timeout = 5 * time.Second
		} else {
			timeout = time.Duration(ctx.keepalive*2) * time.Second

			if timeout < leastReadTimeout {
				timeout = leastReadTimeout
			}
		}

		ctx.Conn.SetReadDeadline(time.Now().Add(timeout))
		if err = de.Read(ctx.Conn); err != nil {
			ctx.exitCode = ecReadError
			ctx.err = err
			log.Infof("Read error: %s", err)
			return
		}

		select {
		case <-ctx.StopCh:
			ctx.exitCode = ecServerClose
			return
		default:
		}

		now := time.Now().Unix()
		if ctx.keepalive != 0 && now-ctx.lastReceived >= int64(ctx.keepalive*2) {
			log.Info("current session is timeout.")
			ctx.exitCode = ecSessionTimeout
			ctx.err = errors.New("Session timeout")
			return
		}

		log.Infof("current session, now : %lld, next refresh : %lld", now, ctx.nextRefresh)
		if now >= ctx.nextRefresh {
			ctx.refreshSession()
		}

		for {
			ret, err := de.DecodeMQTT()
			if err != nil {
				log.Errorf("Invalid mqtt protocol: %s", err)
				ctx.exitCode = ecInvalidProtocol
				ctx.err = err
				return
			}
			ctx.lastReceived = time.Now().Unix()

			if ret != nil {
				err := ctx.handlePacket(ret)
				if err != nil {
					ctx.err = err
					return
				}
			} else {
				break
			}
		}
	}
}

func (ctx *ClientCtx) handlePacket(p interface{}) error {
	log.Infof("Receive pkt: %s", p)
	switch pkt := p.(type) {
	case *mqtt.ConnectPacket:
		return ctx.handleConnectPacket(pkt)
	case *mqtt.PublishPacket:
		return ctx.handlePublishPacket(pkt)
	case *mqtt.PubackPacket:
		return ctx.handlePubackPacket(pkt)
	case *mqtt.SubscribePacket:
		return ctx.handleSubscribePacket(pkt)
	case *mqtt.UnsubscribePacket:
		return ctx.handleUnsubPacket(pkt)
	case *mqtt.PingreqPacket:
		return ctx.handlePingreqPacket(pkt)
	case *mqtt.DisconnectPacket:
		return ctx.handleDisconnectPacket(pkt)
	default:
		// Pubrec, Pubrel, Pubcomp, Suback, Unsuback, Pingresp
		log.Infof("do not support packet.")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Protocol volatilation")
	}

}

var sessionPresent = true

func startSpan() (opentracing.Span, context.Context) {
	span, tctx := opentracing.StartSpanFromContext(context.Background(), "call")
	md, ok := metadata.FromContext(tctx)
	if !ok {
		md = make(map[string]string)
	}
	defer span.Finish()
	// inject opentracing textmap into empty context, for tracking
	opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
	tctx = opentracing.ContextWithSpan(tctx, span)
	tctx = metadata.NewContext(tctx, md)
	return span, tctx
}

func (ctx *ClientCtx) handleConnectPacket(p *mqtt.ConnectPacket) error {
	if ctx.hasConnected {
		log.Errorf("Receive 2nd connect packet")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has connected")
	}

	productKey, err := getProductKey(p.Username)

	if err != nil {
		log.Errorf("Username is invalid: %s", p.Username)
		ctx.exitCode = ecBadClientRequest
		return err
	}

	deviceName := string(p.ClientId)
	ctx.productKey = productKey
	ctx.deviceName = deviceName

	code := 0
	// send connect packet to registry manager for verifying connect package
	registerCli := proto.NewRegistryManagerService(util.REGISTER_MANAGER_SVC, util.Ctx.RegisterSvc.Client())
	conMsg := proto.ConnectMessageRequest{
		ClientId:  p.ClientId,
		Username:  p.Username,
		Password:  string(p.Password),
		WillMsg:   string(p.WillMessage),
		WillTopic: p.WillTopic,
	}

	span, tctx := startSpan()
	defer span.Finish()

	rsp, err := registerCli.Registry(tctx, &conMsg)
	if err != nil {
		code = -1
		log.Errorf("register failed, error message : %s", err.Error())
	} else {
		log.Infof("register success, guid : %d,  message : %s", rsp.Guid, rsp.Message)
	}

	connack := mqtt.NewConnackPacket(sessionPresent, byte(code))

	// Connack is fixed size, so encode will always succeed
	// encode error check is not necessary
	connack.Encode(&ctx.encoder)

	// Now there is only one goroutine to access ctx, no mutex required.
	e := ctx.encoder.WriteTo(ctx.Conn, 5*time.Second)

	if e != nil {
		log.Infof("Write error: %s", e)
		ctx.exitCode = ecWriteError
		return e
	}

	if code != mqtt.RC_ACCEPTED {
		log.Errorf("Mqtt connect error: %d", code)
		ctx.exitCode = ecBadClientRequest
		return errors.New("Connect error")
	}
	ctx.hasConnected = true
	ctx.keepalive = int(p.Keepalive)
	ctx.sessionValue = generateValue(Global.SessionPrefix, ctx.Fd)
	ctx.sessionKey = generateKey(rsp.Guid)
	ctx.guid = rsp.Guid

	// After activate, ctx may be accessed by multi goroutines
	ctx.Mux.Lock()
	ctx.status = activeState
	ctx.Mux.Unlock()

	ctx.refreshSession()
	UpdateDeviceInfo(rsp.Guid, DeviceInfo{
		DeviceId:   deviceName,
		ProductKey: productKey,
	})
	UpdateGuidToDeviceIdMap(rsp.Guid, deviceName)
	return nil
}

const clientSessionKeyNamespace = "C"

func generateKey(guid uint32) string {
	return fmt.Sprintf("%s:%d", clientSessionKeyNamespace, guid)
}

func generateValue(prefix string, fd int) string {
	return fmt.Sprintf("%s:%d", prefix, fd)
}

func getProductKey(username string) (string, error) {
	ss := strings.Split(username, ":")
	if len(ss) == 1 {
		return username, nil
	} else if len(ss) == 2 {
		return ss[0], nil
	}

	return "", errors.New("Invalid username")
}

func refreshDelay() int64 {
	return int64(Global.RedisSessionRefresh/time.Second) + int64(rand.Int()%randRange)
}

const randRange = 4 * 60 // 4 minutes

func (ctx *ClientCtx) refreshSession() {
	log.Info("begin to refresh session.")
	ctx.Mux.RLock()
	if ctx.status != activeState {
		log.Info("current client is not active.")
		ctx.Mux.RUnlock()
		return
	}

	sessions := make(map[string]string, 1)
	sessVal := ctx.sessionValue
	if ctx.productKey != "" {
		sessions[ctx.productKey] = ctx.deviceName
	}
	ctx.Mux.RUnlock()

	sessionStorage := Global.SessionStorage

	for productKey, deviceName := range sessions {
		k := generateKey(ctx.guid)
		oldSessionValue, err := sessionStorage.Refresh(k, sessVal, Global.RedisSessionTimeOut)
		log.Infof("Refresh session: %s, %s, %s, %s", productKey, deviceName, k, sessVal)
		if err != nil {
			log.Errorf("Session refresh error: %v, %s, %s", err, k, sessVal)
		} else if oldSessionValue != "" {
			log.Errorf("Session duplication, kick former, %s, %s, former: %s", k, sessVal, oldSessionValue)
			//kickDupConnection(sessVal, productKey, deviceName, oldSessionValue)
		}
	}
	ctx.nextRefresh = time.Now().Unix() + refreshDelay()
}

func (ctx *ClientCtx) publishForward(p *mqtt.PublishPacket) error {
	if p.Header.Qos() == 2 {
		log.Errorf("Qos 2 is not supported")
		return errors.New("Qos 2 is not supported")
	} else if p.Header.Qos() == 1 {
		span, tctx := startSpan()
		defer span.Finish()

		pubEnginCli := proto.NewPublishEngineService(util.PUBLISH_ENGINE_SVC, util.Ctx.PubEngineSvc.Client())
		pubMsg := proto.PublishMessageRequest{
			Header: &proto.MessageHeader{
				Qos: int32(p.Header.Qos()),
			},
			Topic:   p.Topic,
			Payload: p.Payload,
		}
		_, err := pubEnginCli.PublishMessage(tctx, &pubMsg)
		if err != nil {
			log.Errorf("publish to publish engine failed, error message : %v.", err)
		}

		puback := mqtt.NewPubackPacket(p.PacketId)

		ctx.Mux.Lock()

		// puback is fixed size
		// encode error check is not necessary
		puback.Encode(&ctx.encoder)
		e := ctx.encoder.WriteTo(ctx.Conn, 5*time.Second)

		ctx.Mux.Unlock()

		if e != nil {
			log.Infof("Write error: %s", e)
			ctx.exitCode = ecWriteError
			return e
		}
	}

	return nil
}

func (ctx *ClientCtx) handlePublishPacket(p *mqtt.PublishPacket) error {
	if !ctx.hasConnected {
		log.Error("Publish before connect")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has not been connected before publish")
	}

	return ctx.publishForward(p)
}

func (ctx *ClientCtx) handlePubackPacket(p *mqtt.PubackPacket) error {
	if !ctx.hasConnected {
		log.Error("Receive puback before connect")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has not been connected before")
	}
	log.Infof("Push ack packet ok, packetId: %d", p.PacketId)
	ctx.AckPacket(p.PacketId, rcRequestOK)

	return nil
}

func (ctx *ClientCtx) handleSubscribePacket(p *mqtt.SubscribePacket) error {
	if !ctx.hasConnected {
		log.Error("Receive subscribe before connect")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has not been connected before")
	}

	codes := make([]int8, len(p.Topics))

	span, tctx := startSpan()
	defer span.Finish()

	messageDispatcherCli := proto.NewMessageDispatcherService(util.MESSAGE_DISPATCHER_SVC,
		util.Ctx.MessageDispatcherSvc.Client())
	for i, _ := range p.Topics {
		subMsg := proto.SubscribeMessageRequest{
			TopicFilter: p.Topics[i],
			Qos:         int32(p.Qoss[i]),
			Guid:        ctx.guid,
		}
		span.SetTag("req", subMsg)
		span.SetTag("from-to", "hub-dispater")
		_, err := messageDispatcherCli.Subscribe(tctx, &subMsg)
		if err != nil {
			log.Error("send subscribe message to message dispatcher failed, error message.")
			codes[i] = -2
		} else {
			codes[i] = 1
		}
	}

	ack := mqtt.NewSubackPacket(p.PacketId, codes)

	ctx.Mux.Lock()

	// we limit max topic num per sub is 16
	// so suback payload is small, encode will always succeed
	ack.Encode(&ctx.encoder)
	e := ctx.encoder.WriteTo(ctx.Conn, 5*time.Second)

	ctx.Mux.Unlock()

	log.Info("subscribe success.")
	if e != nil {
		log.Infof("Write error: %s", e)
		ctx.exitCode = ecWriteError
		return e
	}
	return nil
}

func (ctx *ClientCtx) handleUnsubPacket(p *mqtt.UnsubscribePacket) error {
	if !ctx.hasConnected {
		log.Error("Receive unsub before connect")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has not been connected before")
	}

	span, tctx := startSpan()
	defer span.Finish()

	messageDispatcherCli := proto.NewMessageDispatcherService(util.MESSAGE_DISPATCHER_SVC,
		util.Ctx.MessageDispatcherSvc.Client())
	for i, _ := range p.Topics {
		unSubMsg := proto.UnSubscribeMessageRequest{
			TopicFilter: p.Topics[i],
		}
		rsp, err := messageDispatcherCli.UnSubscribe(tctx, &unSubMsg)
		if err != nil {
			log.Errorf("send unsubscribe message to message dispather failed, error message : %s", rsp.Message)
		}
	}

	var e = errors.New("")
	ack := mqtt.NewUnsubackPacket(p.PacketId)

	ctx.Mux.Lock()
	// unsuback is fixed size and small,
	// encode will always succeed
	ack.Encode(&ctx.encoder)
	e = ctx.encoder.WriteTo(ctx.Conn, 5*time.Second)
	ctx.Mux.Unlock()

	if e != nil {
		log.Infof("Write error: %s", e)
		ctx.exitCode = ecWriteError
		return e
	}

	return nil
}

func (ctx *ClientCtx) handlePingreqPacket(p *mqtt.PingreqPacket) error {
	pingresp := mqtt.NewPingrespPacket()

	ctx.Mux.Lock()

	pingresp.Encode(&ctx.encoder)
	e := ctx.encoder.WriteTo(ctx.Conn, 5*time.Second)

	ctx.Mux.Unlock()

	if e != nil {
		log.Infof("Write error: %s", e)
		ctx.exitCode = ecWriteError
		return e
	}
	return nil
}

func (ctx *ClientCtx) handleDisconnectPacket(p *mqtt.DisconnectPacket) error {
	ctx.exitCode = ecClientClose
	return errors.New("Client close the connection")
}

const defaultDup = false
const defaultRetain = false
