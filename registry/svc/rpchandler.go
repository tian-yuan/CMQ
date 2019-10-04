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
	rsp.Code = 200
	rsp.Message = "registry success."
	return nil
}



