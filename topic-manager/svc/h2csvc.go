package svc

import (
	"sync"
	"net"
	"github.com/sirupsen/logrus"
	"strconv"
	"net/http"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type H2cSvc struct {
	Conf *H2cConf

	w *sync.WaitGroup

	tcp net.Listener

	StopCh chan struct{}
}

type H2cConf struct {
	Host string
	Port uint16
}

func NewH2cConf() *H2cConf {
	return &H2cConf{
		Host: "0.0.0.0",
		Port: 1883,
	}
}

func NewH2cSvc(conf *H2cConf) *H2cSvc {
	return &H2cSvc {
		Conf: conf,
		StopCh:  make(chan struct{}),
	}
}

func (cs *H2cSvc) Start() {
	logrus.WithFields(logrus.Fields{
		"Host": cs.Conf.Host,
		"Port": cs.Conf.Port,
	}).Info("start h2c server.")


	addr := net.JoinHostPort(cs.Conf.Host, strconv.Itoa(int(cs.Conf.Port)))
	logrus.Infof("h2c addr: %s", addr)
	h2s := &http2.Server{}

	handler := http.HandlerFunc(handleRequest)

	server := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(handler, h2s),
	}

	server.ListenAndServe()
}
