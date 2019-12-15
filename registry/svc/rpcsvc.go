package svc

import (
	"fmt"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/zookeeper"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/tian-yuan/iot-common/plugins/tracer"
	"github.com/tian-yuan/iot-common/util"
)

type RpcSvc struct {
}

func NewRpcSvc() *RpcSvc {
	return &RpcSvc{}
}

func (svc *RpcSvc) Start(zkAddr []string, tracerAddr string) {
	optFunc := func(opt *registry.Options) {
		opt = &registry.Options{
			Addrs: zkAddr,
		}
	}
	registry := zookeeper.NewRegistry(optFunc)

	t, io, err := tracer.NewTracer(util.REGISTER_MANAGER_SVC, tracerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	// set var t to Global Tracer (opentracing single instance mode)
	opentracing.SetGlobalTracer(t)

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(util.REGISTER_MANAGER_SVC),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())), // add tracing plugin in to middleware
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterRegistryManagerHandler(service.Server(), new(rpchandler))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
