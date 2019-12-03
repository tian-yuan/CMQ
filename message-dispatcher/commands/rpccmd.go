package commands

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"

	"github.com/tian-yuan/CMQ/message-dispatcher/svc"
	"strings"
	"github.com/tian-yuan/iot-common/util"
)

var httpCmd = &cobra.Command{
	Use:   "rpc",
	Short: "start rpc server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start rpc server message dispatcher v0.0.1 -- HEAD")

		var zkAddr string
		cmd.Flags().StringVarP(&zkAddr, "zkAddress", "z", "127.0.0.1:2181", "zk address array")
		logrus.Infof("start discovery client, zk address : %s", zkAddr)
		zkAddrArr := strings.Split(zkAddr, ";")
		util.Ctx.InitTopicManagerSvc(zkAddrArr)

		topicSvc := svc.NewTopicLoadSvc()
		svc.Global.TopicLoadSvc = topicSvc
		go topicSvc.Start(zkAddrArr)

		rpcSvc := svc.NewRpcSvc()
		util.Ctx.InitTopicAclSvc(zkAddrArr)
		rpcSvc.Start(zkAddrArr)
	},
}
