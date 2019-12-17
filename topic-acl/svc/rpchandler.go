package svc

import (
	"context"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/util/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	proto "github.com/tian-yuan/iot-common/iotpb"
)

type rpchandler struct {
}

func (r *rpchandler) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	// get tracing info from context
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	sp = opentracing.StartSpan("Subscribe", opentracing.ChildOf(wireContext))
	// record request
	sp.SetTag("req", in)
	defer func() {
		// record response
		sp.SetTag("res", out)
		// before function return stop span, cuz span will counted how much time of this function spent
		sp.Finish()
	}()

	if in.Guid < 0 ||
		in.Qos < 0 ||
		in.TopicFilter == "" {
		out.Code = 430
		out.Message = "parameter error."
		log.Error("parameter error.")
		return errors.New("parameter error.")
	}

	topicId, err := Global.TopicSvc.Subscribe(in.TopicFilter, in.Guid, int(in.Qos))
	out.TopicId = topicId
	if err != nil {
		out.Code = 600
		out.Message = err.Error()
	}
	return nil
}

func (r *rpchandler) UnSubscribe(ctx context.Context, in *proto.UnSubscribeMessageRequest, out *proto.UnSubscribeMessageResponse) error {
	// get tracing info from context
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	sp = opentracing.StartSpan("Subscribe", opentracing.ChildOf(wireContext))
	// record request
	sp.SetTag("req", in)
	defer func() {
		// record response
		sp.SetTag("res", out)
		// before function return stop span, cuz span will counted how much time of this function spent
		sp.Finish()
	}()

	if in.Guid < 0 ||
		in.TopicFilter == "" {
		out.Code = 430
		out.Message = "parameter error."
		log.Error("parameter error.")
		return errors.New("parameter error.")
	}
	_, err := Global.TopicSvc.UnSubscribe(in.TopicFilter, in.Guid)
	if err != nil {
		out.Code = 600
		out.Message = err.Error()
	}
	return nil
}
