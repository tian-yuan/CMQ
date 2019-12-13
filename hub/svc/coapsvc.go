package svc

import (
	"net"
	"strconv"
	"sync"

	"github.com/micro/go-micro/util/log"
	"github.com/tian-yuan/CMQ/hub/proto/coap"
)

type CoapSvc struct {
	Conf *CoapConf

	w *sync.WaitGroup

	tcp net.Listener

	StopCh chan struct{}
}

type CoapConf struct {
	Type     string
	CoapHost string
	CoapPort uint16
}

func NewCoapConf() *CoapConf {
	return &CoapConf{
		Type:     "udp",
		CoapHost: "0.0.0.0",
		CoapPort: 1883,
	}
}

func NewCoapSvc(conf *CoapConf) *CoapSvc {
	return &CoapSvc{
		Conf:   conf,
		StopCh: make(chan struct{}),
	}
}

func (cs *CoapSvc) Start() {
	log.Info("start coap server.")

	mux := coap.NewServeMux()

	coapaddr := net.JoinHostPort(cs.Conf.CoapHost, strconv.Itoa(int(cs.Conf.CoapPort)))
	log.Infof("coap addr: %s", coapaddr)
	go coap.ListenAndServe(cs.Conf.Type, coapaddr, mux)
}
