package svc

import (
	"context"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/sirupsen/logrus"
	"github.com/tian-yuan/iot-common/util"
)

type rpchandler struct {

}

func (h *rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	logrus.Infof("publish engine service receive publish message, topic: %s, payload : %s",
		in.Topic, string(in.Payload))
	messageDispatcherCli := proto.NewMessageDispatcherService(util.MESSAGE_DISPATCHER_SVC,
		util.Ctx.MessageDispatcherSvc.Client())
	out, err := messageDispatcherCli.PublishMessage(context.TODO(), in)
	if err != nil {
		logrus.Error("publish to message dispatcher failed.")
	} else {
		logrus.Infof("publish to message dispatcher success, rsp : %s", out.Message)
	}
	return nil
}

