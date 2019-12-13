package svc

import (
	"net"
	"strconv"
	"sync"

	"github.com/micro/go-micro/util/log"
	"github.com/tian-yuan/iot-common/util"
)

type MqttSvc struct {
	Conf *MqttConf

	w *sync.WaitGroup

	tcp net.Listener

	StopCh chan struct{}
}

type MqttConf struct {
	MqttHost string
	MqttPort uint16
}

func init() {
	ClientCtxs = make([]ClientCtx, 50*10000)
	for i := 0; i < len(ClientCtxs); i += 1 {
		ClientCtxs[i].Fd = i
	}
}

func NewMqttConf() *MqttConf {
	return &MqttConf{
		MqttHost: "0.0.0.0",
		MqttPort: 1883,
	}
}

func NewMqttSvc(conf *MqttConf) *MqttSvc {
	return &MqttSvc{
		Conf:   conf,
		StopCh: make(chan struct{}),
	}
}

func (ms *MqttSvc) Start() {
	log.Info("start mqtt server.")

	tcpaddr := net.JoinHostPort(ms.Conf.MqttHost, strconv.Itoa(int(ms.Conf.MqttPort)))
	log.Infof("Tcp addr: %s", tcpaddr)
	tcpl, e := util.NewTCPListener(tcpaddr, false)
	if e != nil {
		log.Fatalf("New tcp listener error: %v", e)
		return
	}

	ms.tcp = tcpl
	go startListen(ms, tcpl)
	log.Info("Start to listen tcp...")
	return
}

func startListen(s *MqttSvc, l net.Listener) {
	for {
		select {
		case <-s.StopCh:
			return
		default:
		}

		c, err := l.Accept()
		if err == nil {
			go handleConnection(s, c, false)
		} else {
			select {
			case <-s.StopCh:
				return
			default:
			}

			log.Errorf("Accept error, listener: %+v, error: %v", l, err)
		}
	}
}

func (ms *MqttSvc) Stop() {

}
