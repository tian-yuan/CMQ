package svc

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/util/log"
	opentracing "github.com/opentracing/opentracing-go"
	proto "github.com/tian-yuan/iot-common/iotpb"
)

type rpchandler struct {
}

const hubPublishPath = "/v1/imd/publish"

func publishMessageToHub(ctx context.Context, info string, message string, topic string, qos int8) error {
	log.Infof("publish message to hub, info : %s, message : %s, topic : %s, qos : %d",
		info, message, topic, qos)
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	span := opentracing.StartSpan("PublishMessage-IotHub", opentracing.ChildOf(wireContext))
	// record request
	// before function return stop span, cuz span will counted how much time of this function spent
	defer span.Finish()

	span.SetTag("hub-addr", info)
	// split session info
	infos := strings.Split(info, ":")
	addr := infos[0]
	port := infos[1]
	fd := infos[2]
	cli := GetH2client(addr + ":" + port)
	u := &url.URL{
		Scheme: "http",
		Host:   addr + ":" + port,
		Path:   hubPublishPath,
	}

	query := url.Values{}
	query.Add("fd", fd)
	query.Add("message", message)
	query.Add("topic", topic)
	query.Add("qos", fmt.Sprintf("%d", qos))
	u.RawQuery = query.Encode()
	log.Infof("publish message to hub : %s", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Errorf("Construct http request error")
		span.SetTag("Message", err.Error())
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	_, err = cli.Do(req)
	if err != nil {
		log.Errorf("publish message to hub failed.")
		span.SetTag("Message", err.Error())
		return nil
	}
	log.Infof("publish message to hub success.")
	return nil
}

func (h *rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	log.Infof("topic manager receive publish message.")
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	sp = opentracing.StartSpan("PublishMessageMatch", opentracing.ChildOf(wireContext))
	// record request
	sp.SetTag("req", in)
	sp.SetTag("res", out)
	sp.Finish()

	subs := Ctx.Match(in.Topic)
	sp.SetTag("subs", subs)
	for _, sub := range subs {
		// get guid from subs
		guid := sub.(uint32)
		log.Infof("device guid match topic : %s, guid : %d", in.Topic, guid)
		// get session info from redis by guid
		r := Global.RedisClient.Get("C:" + fmt.Sprintf("%d", guid))
		if err := r.Err(); err != nil {
			log.Error(err, "get session info from redis failed, guid : "+string(guid))
			continue
		}
		// send request to hub
		val := r.Val()
		publishMessageToHub(ctx, val, string(in.Payload), in.Topic, int8(in.Header.Qos))
	}
	out.Code = 200
	out.Message = "publish success."
	return nil
}

func (h *rpchandler) LoadSubTopic(ctx context.Context, in *proto.SubTopicLoadRequest, out *proto.SubTopicLoadResponse) error {
	log.Infof("handle load sub topic request, product key : %s.", in.ProductKey)
	return nil
}

func (h *rpchandler) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	log.Infof("handler subscribe message, topic : %s", in.TopicFilter)
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

	err := Ctx.Subscribe(in.TopicFilter, int(in.Qos), in.Guid)
	if err != nil {
		log.Errorf("subscribe topic failed, topic : %s", in.TopicFilter)
		out.Code = 600
		out.Message = err.Error()
	}
	// should write to database
	out.Code = 200
	out.Message = "subscribe success."
	log.Infof("handler subscribe message success, topic : %s", in.TopicFilter)
	return nil
}

func (h *rpchandler) UnSubscribe(ctx context.Context, in *proto.UnSubscribeMessageRequest, out *proto.UnSubscribeMessageResponse) error {
	log.Info("handle unsubscribe message request.")
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

	return nil
}
