package commands

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/micro/go-micro/util/log"
	"github.com/tian-yuan/CMQ/message-dispatcher/svc"
	"github.com/tian-yuan/iot-common/util"
)

var httpCmd = &cobra.Command{
	Use:   "rpc",
	Short: "start rpc server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Start rpc server message dispatcher v0.0.1 -- HEAD")

		var etcdAddr string
		cmd.Flags().StringVarP(&etcdAddr, "etcdAddress", "e", "127.0.0.1:2379", "etcd address array")
		etcdAddrArr := strings.Split(etcdAddr, ";")
		log.Infof("start discovery client, etcd address : %s", etcdAddr)

		var zkAddr string
		cmd.Flags().StringVarP(&zkAddr, "zkAddress", "z", "127.0.0.1:2181", "zookeeper address array")
		zkAddrArr := strings.Split(zkAddr, ";")
		log.Infof("start discovery client, zookeeper address : %s", zkAddr)

		var tracerAddr string
		cmd.Flags().StringVarP(&tracerAddr, "tracerAddress", "j", "127.0.0.1:6831", "tracer address array")

		util.Init(
			util.WithRegistryUrls(etcdAddr),
			util.WithTracerUrl(tracerAddr),
		)

		util.Ctx.InitTopicAclSvc()
		defer util.Ctx.CloseTopicAclSvc()
		util.Ctx.InitTopicManagerSvc()
		defer util.Ctx.CloseTopicManagerSvc()

		topicSvc := svc.NewTopicLoadSvc()
		svc.Global.TopicLoadSvc = topicSvc
		go topicSvc.Start(zkAddrArr, etcdAddrArr, tracerAddr)

		rpcSvc := svc.NewRpcSvc()
		rpcSvc.Start(etcdAddrArr, tracerAddr)
	},
}
