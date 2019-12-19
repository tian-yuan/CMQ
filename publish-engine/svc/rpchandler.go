package svc

import (
	"context"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/util/log"
	opentracing "github.com/opentracing/opentracing-go"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/tian-yuan/iot-common/util"
)

type rpchandler struct {
}

func (h *rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	log.Infof("publish engine service receive publish message, topic: %s, payload : %s",
		in.Topic, string(in.Payload))

	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	ssp := opentracing.StartSpan("PublishMessage", opentracing.ChildOf(wireContext))
	// record request
	ssp.SetTag("req", in)
	// before function return stop span, cuz span will counted how much time of this function spent
	defer ssp.Finish()

	ssp.SetTag("req", in)
	messageDispatcherCli := proto.NewMessageDispatcherService(util.MESSAGE_DISPATCHER_SVC,
		util.Ctx.MessageDispatcherSvc.Client())
	out, err := messageDispatcherCli.PublishMessage(opentracing.ContextWithSpan(context.Background(), ssp), in)
	if err != nil {
		log.Error("publish to message dispatcher failed.")
	} else {
		log.Infof("publish to message dispatcher success, rsp : %s", out.Message)
	}
	return nil
}
