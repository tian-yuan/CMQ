package svc

import (
	"sync"
	"net"
	"github.com/sirupsen/logrus"
	"CMQ/hub/proto/coap"
	"strconv"
)

type CoapSvc struct {
	Conf *CoapConf

	w *sync.WaitGroup

	tcp net.Listener

	StopCh chan struct{}
}

type CoapConf struct {
	Type string
	CoapHost string
	CoapPort uint16
}

func NewCoapConf() *CoapConf {
	return &CoapConf{
		Type: "udp",
		CoapHost: "0.0.0.0",
		CoapPort: 1883,
	}
}

func NewCoapSvc(conf *CoapConf) *CoapSvc {
	return &CoapSvc {
		Conf: conf,
		StopCh:  make(chan struct{}),
	}
}

func (cs *CoapSvc) Start() {
	logrus.WithFields(logrus.Fields{
		"type": cs.Conf.Type,
		"coapHost": cs.Conf.CoapHost,
		"coapPort": cs.Conf.CoapPort,
	}).Info("start coap server.")

	mux := coap.NewServeMux()

	coapaddr := net.JoinHostPort(cs.Conf.CoapHost, strconv.Itoa(int(cs.Conf.CoapPort)))
	logrus.Infof("coap addr: %s", coapaddr)
	go coap.ListenAndServe(cs.Conf.Type, coapaddr, mux)
}
