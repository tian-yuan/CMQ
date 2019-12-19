package svc

import (
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/micro/go-micro/util/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/tian-yuan/iot-common/plugins/tracer"
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
	return &H2cSvc{
		Conf:   conf,
		StopCh: make(chan struct{}),
	}
}

func (cs *H2cSvc) Start(tracerAddr string) {
	log.Info("start h2c server.")

	addr := net.JoinHostPort(cs.Conf.Host, strconv.Itoa(int(cs.Conf.Port)))
	log.Infof("h2c addr: %s", addr)
	h2s := &http2.Server{}

	handler := http.HandlerFunc(handleRequest)

	server := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(handler, h2s),
	}

	t, io, err := tracer.NewTracer("iot.hub.svc", tracerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	server.ListenAndServe()
}
