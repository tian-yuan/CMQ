package svc

import (
	"context"

	"github.com/micro/go-micro/util/log"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/tian-yuan/iot-common/util"
)

type rpchandler struct {
}

func (h *rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	log.Infof("publish engine service receive publish message, topic: %s, payload : %s",
		in.Topic, string(in.Payload))
	messageDispatcherCli := proto.NewMessageDispatcherService(util.MESSAGE_DISPATCHER_SVC,
		util.Ctx.MessageDispatcherSvc.Client())
	out, err := messageDispatcherCli.PublishMessage(context.TODO(), in)
	if err != nil {
		log.Error("publish to message dispatcher failed.")
	} else {
		log.Infof("publish to message dispatcher success, rsp : %s", out.Message)
	}
	return nil
}
