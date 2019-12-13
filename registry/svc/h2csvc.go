package svc

import (
	"net"
	"net/http"
	"strconv"

	"github.com/micro/go-micro/util/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type H2cSvc struct {
	Conf *H2cConf
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
	return &H2cSvc{
		Conf: conf,
	}
}

func (cs *H2cSvc) Start() {
	log.Info("start h2c server.")

	addr := net.JoinHostPort(cs.Conf.Host, strconv.Itoa(int(cs.Conf.Port)))
	log.Infof("h2c addr: %s", addr)
	h2s := &http2.Server{}

	handler := http.HandlerFunc(handleRequest)

	server := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(handler, h2s),
	}

	server.ListenAndServe()
}
