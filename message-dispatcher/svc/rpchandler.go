package svc

import (
	"context"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/sirupsen/logrus"
)

type rpchandler struct {

}

func (h *rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	logrus.Infof("message dispatcher service receive publish message, topic: %s, payload : %s",
		in.Topic, string(in.Payload))
	return nil
}

func (h *rpchandler) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	return nil
}

func (h *rpchandler) UnSubscribe(ctx context.Context, in *proto.UnSubscribeMessageRequest, out *proto.UnSubscribeMessageResponse) error {
	return nil
}