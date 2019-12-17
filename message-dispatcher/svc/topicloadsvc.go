package svc

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/zookeeper"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/tian-yuan/CMQ/message-dispatcher/consistent"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/tian-yuan/iot-common/plugins/tracer"
	"github.com/tian-yuan/iot-common/util"
)

type TopicLoadSvc struct {
	ServiceCache ServiceCache

	Consistent *consistent.Consistent

	client client.Client
	closer io.Closer
	done   chan bool
	ticker *time.Ticker

	sync.RWMutex
	lock *zk.Lock
}

func NewTopicLoadSvc() *TopicLoadSvc {
	return &TopicLoadSvc{}
}

func (svc *TopicLoadSvc) startTopicController(zkAddr []string) {
	for {
		cli, _, err := zk.Connect(zkAddr, time.Second*30)
		if err != nil {
			time.Sleep(time.Second)
			log.Errorf("connect to zk failed : %v", err.Error())
			continue
		}
		acls := zk.WorldACL(zk.PermAll)
		svc.lock = zk.NewLock(cli, "/iot-message-dispatcher-controller", acls)
		err = svc.lock.Lock()
		if err != nil && err == zk.ErrDeadlock {
			log.Info("already get controller lock.")
			break
		} else if err != nil && err != zk.ErrDeadlock {
			time.Sleep(time.Second)
			log.Errorf("create lock error : %v", err.Error())
			continue
		} else {
			log.Info("get controller lock success.")
			break
		}
	}
}

func (svc *TopicLoadSvc) Start(zkAddr []string, tracerAddr string) {
	svc.startTopicController(zkAddr)

	svc.Consistent = consistent.New()

	optFunc := func(opt *registry.Options) {
		opt = &registry.Options{
			Addrs: zkAddr,
		}
	}

	r := zookeeper.NewRegistry(optFunc)
	c := client.NewClient(
		client.Registry(r),
	)
	c.Options().Selector.Init(selector.Registry(r))

	clientSvcName := util.TOPIC_MANAGER_SVC + ".client"
	t, io, err := tracer.NewTracer(clientSvcName, tracerAddr)
	if err == nil {
		clientWrapper := ocplugin.NewClientWrapper(t)
		svc.client = clientWrapper(c)
		svc.closer = io
	} else {
		svc.client = c
	}

	svc.ServiceCache = New(r, handleServiceUpdate)

	svc.ServiceCache.GetService(util.TOPIC_MANAGER_SVC)

	svc.done = make(chan bool)
	svc.ticker = time.NewTicker(10000 * time.Millisecond)
	go func() {
		for {
			select {
			case <-svc.done:
				return
			case <-svc.ticker.C:
				Global.TopicLoadSvc.ServiceCache.GetService(util.TOPIC_MANAGER_SVC)
			}
		}
	}()
}

func (svc *TopicLoadSvc) Stop() {
	if svc.closer != nil {
		svc.closer.Close()
	}
	svc.ticker.Stop()
	svc.done <- true
	svc.lock.Unlock()
}

func (svc *TopicLoadSvc) TopicReloadRequest(address string, seg string) error {
	log.Infof("service address : %s, should load seg : %s", address, seg)
	service := util.TOPIC_MANAGER_SVC
	endpoint := "TopicManager.LoadSubTopic"
	reqInstance := &proto.SubTopicLoadRequest{
		ProductKey: "test" + address,
	}

	req := svc.client.NewRequest(service, endpoint, reqInstance)

	if err := svc.client.Call(context.Background(), req, nil, client.WithAddress(address)); err != nil {
		log.Errorf("call with address error", err)
		return err
	}
	return nil
}

func (svc *TopicLoadSvc) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	// should send to topic acl and get topic id
	// should send topic subscribe request to acl to check the acl and obtain the topic id
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	ssp := opentracing.StartSpan("Subscribe-TopicAcl", opentracing.ChildOf(wireContext))
	// record request
	ssp.SetTag("req", in)
	// before function return stop span, cuz span will counted how much time of this function spent
	defer ssp.Finish()

	topicAclCli := proto.NewTopicAclService(util.TOPIC_ACL_SVC,
		util.Ctx.TopicAclSvc.Client())
	out, err := topicAclCli.Subscribe(opentracing.ContextWithSpan(context.Background(), ssp), in)
	if err != nil {
		log.Error("publish subscribe message to topic acl service failed.")
		return err
	} else {
		log.Infof("publish subscribe message to topic acl success, rsp : %s", out.Message)
	}

	tssp := opentracing.StartSpan("Subscribe-TopicManager", opentracing.ChildOf(wireContext))
	defer tssp.Finish()
	addr, err := Global.TopicLoadSvc.Consistent.Get(fmt.Sprintf("%d", out.TopicId))
	tssp.SetTag("req", in)
	tssp.SetTag("topic.load.svc.addr", addr)
	if err != nil || addr == "" {
		log.Errorf("get topic manager service failed.")
		return err
	}
	service := util.TOPIC_MANAGER_SVC
	endpoint := "TopicManager.Subscribe"

	req := svc.client.NewRequest(service, endpoint, in)

	if err := svc.client.Call(opentracing.ContextWithSpan(context.Background(), tssp), req, out, client.WithAddress(addr)); err != nil {
		log.Errorf("call with address error", err)
		return err
	}
	return nil
}

func (svc *TopicLoadSvc) UnSubscribe(in *proto.UnSubscribeMessageRequest, out *proto.UnSubscribeMessageResponse) error {
	// should send to topic acl and get topic id
	addr, err := Global.TopicLoadSvc.Consistent.Get(fmt.Sprintf("%d", in.Guid))
	if err != nil {
		log.Errorf("get topic manager service failed.")
		return err
	}
	service := util.TOPIC_MANAGER_SVC
	endpoint := "TopicManager.UnSubscribe"

	req := svc.client.NewRequest(service, endpoint, in)

	if err := svc.client.Call(context.Background(), req, out, client.WithAddress(addr)); err != nil {
		log.Errorf("call with address error", err)
		return err
	}
	return nil
}

func (svc *TopicLoadSvc) PublishMessage(in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	services, err := Global.TopicLoadSvc.ServiceCache.GetService(util.TOPIC_MANAGER_SVC)
	if err != nil {
		log.Errorf("get service for topic manager failed.", err)
		return err
	}

	if services == nil || len(services) == 0 {
		return fmt.Errorf("get service empty.")
	}

	for _, service := range services {
		for _, node := range service.Nodes {
			service := util.TOPIC_MANAGER_SVC
			endpoint := "TopicManager.PublishMessage"

			req := svc.client.NewRequest(service, endpoint, in)

			if err := svc.client.Call(context.Background(), req, out, client.WithAddress(node.Address)); err != nil {
				log.Errorf("call with address error", err)
				// if qos is 1 or 2, we should add the message to queue
			}
		}
	}
	return nil
}

func handleServiceUpdate(service string, services []*registry.Service) {
	log.Infof("handle service :%s update", service)
	if service != util.TOPIC_MANAGER_SVC {
		log.Infof("unkown service name : %s", service)
		return
	}

	go topicReload(services)
}

const TOPIC_ID_RANGE = 20000

func topicReload(services []*registry.Service) {
	if services == nil || len(services) == 0 {
		log.Info("all service is down.")
		return
	}

	if services == nil {
		return
	}
	c := consistent.New()
	serviceMap := make(map[string]string)
	for _, service := range services {
		for _, node := range service.Nodes {
			c.Add(node.Address)
			serviceMap[node.Address] = ""
		}
	}
	// get max topic id
	var maxTopicId uint32
	maxTopicId = 20*10000 + 1
	// send request to topic manager service
	SegIdRange := maxTopicId / TOPIC_ID_RANGE
	if (SegIdRange % TOPIC_ID_RANGE) > 0 {
		SegIdRange += 1
	}
	var i uint32
	for i = 0; i < SegIdRange; i++ {
		addr, err := c.Get(fmt.Sprintf("%d", i))
		if err == nil && addr != "" {
			serviceMap[addr] += fmt.Sprintf(",%d", i)
		}
	}
	for addr, seg := range serviceMap {
		log.Infof("service %s should load topic seg info : %s", addr, seg)
		Global.TopicLoadSvc.TopicReloadRequest(addr, seg)
	}

	Global.TopicLoadSvc.Lock()
	defer Global.TopicLoadSvc.Unlock()
	Global.TopicLoadSvc.Consistent = c
}
