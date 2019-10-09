package svc

import (
	"context"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/sirupsen/logrus"
)

type rpchandler struct {

}

func (h* rpchandler) Registry(ctx context.Context, req *proto.ConnectMessageRequest, rsp *proto.ConnectMessageResponse) error {
	logrus.Infof("registry connect message, username : %s, client id : %s", req.Username, req.ClientId)
	// 1. update device info to mysql database and get guid from mysql database
	guid := 1
	rsp.Guid = uint32(guid)
	rsp.Code = 200
	rsp.Message = "registry success."
	return nil
}



