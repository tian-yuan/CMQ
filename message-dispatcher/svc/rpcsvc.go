package svc

import (
	"github.com/micro/go-micro"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"fmt"
)

type RpcSvc struct {

}

func NewRpcSvc() *RpcSvc {
	return &RpcSvc{}
}

func (svc *RpcSvc) Start() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("iot-message-dispatcher"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterMessageDispatcherHandler(service.Server(), new(rpchandler))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
