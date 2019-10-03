package svc

import (
	"github.com/tian-yuan/CMQ/util"
	"net"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"github.com/tian-yuan/CMQ/hub/proto/mqtt"
	"errors"
	topic "github.com/tian-yuan/CMQ/topic-manager/topic"
)

var matcher topic.Matcher

func init() {
	rand.Seed(time.Now().Unix())
	matcher = topic.NewTrieMatcher()
}

func handleConnection(s *MqttSvc, c net.Conn, ws bool) {
	fd := util.GetConnFd(c)

	logrus.Infof("receive connection fd : %d", fd)
	ctx := GetCtxForFd(fd)
	if ctx == nil {
		c.Close()
		return
	}

	logrus.Infof("Client connect, address: %s, isWs: %t, fd: %d", c.RemoteAddr(), ws, fd)

	ctx.Reset(ws, c)

	ctx.handleClient()
	ctx.Mux.Lock()

	logrus.Infof("Client disconnect, address: %s, isWs: %t, fd: %d, reason: %s, err: %s",
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
			logrus.Infof("Read error: %s", err)
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
			ctx.exitCode = ecSessionTimeout
			ctx.err = errors.New("Session timeout")
			return
		}

		if now >= ctx.nextRefresh {
			ctx.refreshSession()
		}

		for {
			ret, err := de.DecodeMQTT()
			if err != nil {
				logrus.Errorf("Invalid mqtt protocol: %s", err)
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
	logrus.Infof("Receive pkt: %s", p)
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
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Protocol volatilation")
	}

}

var sessionPresent = true

func (ctx *ClientCtx) handleConnectPacket(p *mqtt.ConnectPacket) error {
	if ctx.hasConnected {
		logrus.Errorf("Receive 2nd connect packet")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has connected")
	}

	code := 0
	connack := mqtt.NewConnackPacket(sessionPresent, byte(code))

	// Connack is fixed size, so encode will always succeed
	// encode error check is not necessary
	connack.Encode(&ctx.encoder)

	// Now there is only one goroutine to access ctx, no mutex required.
	e := ctx.encoder.WriteTo(ctx.Conn, 5 * time.Second)

	if e != nil {
		logrus.Infof("Write error: %s", e)
		ctx.exitCode = ecWriteError
		return e
	}

	if code != mqtt.RC_ACCEPTED {
		logrus.Errorf("Mqtt connect error: %d", code)
		ctx.exitCode = ecBadClientRequest
		return errors.New("Connect error")
	}
	ctx.hasConnected = true
	ctx.keepalive = int(p.Keepalive)

	// After activate, ctx may be accessed by multi goroutines
	ctx.Mux.Lock()
	ctx.status = activeState
	ctx.Mux.Unlock()

	ctx.refreshSession()
	return nil
}

const clientSessionKeyNamespace = "C"

const randRange = 4 * 60 // 4 minutes

func (ctx *ClientCtx) refreshSession() {

}

func (ctx *ClientCtx) publishForward(p *mqtt.PublishPacket) error {
	if p.Header.Qos() == 2 {
		logrus.Errorf("Qos 2 is not supported")
		return errors.New("Qos 2 is not supported")
	} else if p.Header.Qos() == 1 {
		topic := p.Topic
		subs := matcher.Lookup(topic)
		logrus.Infof("topic : %s match subscriber size : %d", topic, len(subs))
		for _, subFd := range subs {
			// send data to fd
			destCtx := GetCtxForFd(int64(subFd.(int)))
			if destCtx != nil {
				h := mqtt.NewPublishPacketTypeHeader(1, defaultDup, defaultRetain)
				pktId := destCtx.GetPacketId()

				p := mqtt.NewPublishPacket(h, topic, pktId, p.Payload)
				destCtx.publish(p)
			} else {
				logrus.Infof("topic : %s is subscribed by nobody", topic)
			}
		}

		puback := mqtt.NewPubackPacket(p.PacketId)

		ctx.Mux.Lock()

		// puback is fixed size
		// encode error check is not necessary
		puback.Encode(&ctx.encoder)
		e := ctx.encoder.WriteTo(ctx.Conn, 5 * time.Second)

		ctx.Mux.Unlock()

		if e != nil {
			logrus.Infof("Write error: %s", e)
			ctx.exitCode = ecWriteError
			return e
		}
	}

	return nil
}

func (ctx *ClientCtx) handlePublishPacket(p *mqtt.PublishPacket) error {
	if !ctx.hasConnected {
		logrus.Error("Publish before connect")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has not been connected before publish")
	}

	return ctx.publishForward(p)
}

func (ctx *ClientCtx) handlePubackPacket(p *mqtt.PubackPacket) error {
	if !ctx.hasConnected {
		logrus.Error("Receive puback before connect")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has not been connected before")
	}
	logrus.Infof("Push ack packet ok, packetId: %d", p.PacketId)
	ctx.AckPacket(p.PacketId, rcRequestOK)

	return nil
}

func (ctx *ClientCtx) handleSubscribePacket(p *mqtt.SubscribePacket) error {
	if !ctx.hasConnected {
		logrus.Error("Receive subscribe before connect")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has not been connected before")
	}

	codes := make([]int8, 2)
	for _, topic := range p.Topics {
		logrus.Infof("subscribe topic : %s, fd : %d", topic, ctx.Fd)
		matcher.Subscribe(topic, ctx.Fd)
	}

	// if internal call error, set codes to -2
	if codes == nil {
		codes = make([]int8, len(p.Qoss))
		for i, _ := range codes {
			codes[i] = -2
		}
	}

	ack := mqtt.NewSubackPacket(p.PacketId, codes)

	ctx.Mux.Lock()

	// we limit max topic num per sub is 16
	// so suback payload is small, encode will always succeed
	ack.Encode(&ctx.encoder)
	e := ctx.encoder.WriteTo(ctx.Conn, 5 * time.Second)

	ctx.Mux.Unlock()

	if e != nil {
		logrus.Infof("Write error: %s", e)
		ctx.exitCode = ecWriteError
		return e
	}
	return nil
}

func (ctx *ClientCtx) handleUnsubPacket(p *mqtt.UnsubscribePacket) error {
	if !ctx.hasConnected {
		logrus.Error("Receive unsub before connect")
		ctx.exitCode = ecInvalidProtocol
		return errors.New("Client has not been connected before")
	}

	for _, topicName := range p.Topics {
		matcher.Unsubscribe(topic.NewSubscription(0, topicName, ctx.Fd))
	}

	var e = errors.New("")
	ack := mqtt.NewUnsubackPacket(p.PacketId)

	ctx.Mux.Lock()
	// unsuback is fixed size and small,
	// encode will always succeed
	ack.Encode(&ctx.encoder)
	e = ctx.encoder.WriteTo(ctx.Conn, 5 * time.Second)
	ctx.Mux.Unlock()

	if e != nil {
		logrus.Infof("Write error: %s", e)
		ctx.exitCode = ecWriteError
		return e
	}

	return nil
}

func (ctx *ClientCtx) handlePingreqPacket(p *mqtt.PingreqPacket) error {
	pingresp := mqtt.NewPingrespPacket()

	ctx.Mux.Lock()

	pingresp.Encode(&ctx.encoder)
	e := ctx.encoder.WriteTo(ctx.Conn, 5 * time.Second)

	ctx.Mux.Unlock()

	if e != nil {
		logrus.Infof("Write error: %s", e)
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

func (ctx *ClientCtx) publish(p *mqtt.PublishPacket) int {
	if e := p.Encode(&ctx.encoder); e != nil {
		ctx.encoder.ResetState()
		logrus.Errorf("Encode payload error: %v", e)
		return rcEncodeError
	}

	if e := ctx.encoder.WriteTo(ctx.Conn, 5 * time.Second); e != nil {
		logrus.Errorf("Push message error: %v", e)
		return rcNetworkError
	}

	return rcRequestOK
}
