package commands

import (
	"github.com/spf13/cobra"

	"strings"

	"github.com/micro/go-micro/util/log"
	"github.com/tian-yuan/CMQ/topic-acl/svc"
	"github.com/tian-yuan/iot-common/util"
)

var httpCmd = &cobra.Command{
	Use:   "rpc",
	Short: "start rpc server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Start rpc server message dispatcher v0.0.1 -- HEAD")

		var etcdAddr string
		cmd.Flags().StringVarP(&etcdAddr, "etcdAddress", "z", "127.0.0.1:2379", "etcd address array")
		log.Infof("start discovery client, etcd address : %s", etcdAddr)
		etcdAddrArr := strings.Split(etcdAddr, ";")
		var tracerAddr string
		cmd.Flags().StringVarP(&tracerAddr, "tracerAddress", "j", "127.0.0.1:6831", "tracer address array")

		util.Init(
			util.WithRegistryUrls(etcdAddr),
			util.WithTracerUrl(tracerAddr),
		)
		util.Ctx.InitTopicManagerSvc()
		defer util.Ctx.CloseTopicAclSvc()

		log.Info("begin to start mysql client.")
		var mysqlhost string
		var mysqlport uint16
		var username string
		var password string
		var database string
		cmd.Flags().StringVarP(&mysqlhost, "mysqlhost", "m", "127.0.0.1", "mysql hostname")
		cmd.Flags().Uint16VarP(&mysqlport, "mysqlport", "p", 3306, "mysql port")
		cmd.Flags().StringVarP(&username, "username", "u", "root", "mysql username")
		cmd.Flags().StringVarP(&password, "password", "w", "123456", "mysql password")
		cmd.Flags().StringVarP(&database, "database", "b", "push_device", "mysql database")
		conf := svc.NewTopicConfig()
		conf.Host = mysqlhost
		conf.Port = mysqlport
		conf.Username = username
		conf.Password = password
		conf.Database = database
		svc.Global.TopicSvc = svc.NewTopicSvc(conf)
		svc.Global.TopicSvc.Start()

		rpcSvc := svc.NewRpcSvc()
		rpcSvc.Start(etcdAddrArr, tracerAddr)
	},
}
