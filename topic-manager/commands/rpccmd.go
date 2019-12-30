package commands

import (
	"github.com/spf13/cobra"

	"strings"
	"time"

	"github.com/micro/go-micro/util/log"
	"github.com/tian-yuan/CMQ/topic-manager/svc"
	"github.com/tian-yuan/iot-common/util"
)

var rpccmd = &cobra.Command{
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
		defer util.Ctx.CloseTopicManagerSvc()

		var redisClusterAddr string
		cmd.Flags().StringVarP(&redisClusterAddr, "redisClusterAddr", "r", "127.0.0.1:7000", "redis cluster address")
		var reqTimeout time.Duration
		cmd.Flags().DurationVarP(&reqTimeout, "redisReqTimeout", "t", 5*time.Second, "redis request timeout")
		var poolSize int32
		cmd.Flags().Int32VarP(&poolSize, "redisPoolSize", "s", 5, "redis client pool size")
		var redisSessionTimeout time.Duration
		cmd.Flags().DurationVarP(&redisSessionTimeout, "redisSessionTimeout", "o", 20*time.Minute, "redis session timeout")
		var redisSessionRefresh time.Duration
		cmd.Flags().DurationVarP(&redisSessionRefresh, "redisSessionRefresh", "f", 20*time.Minute, "redis session refresh")
		redisClient := util.GetClusterClient(redisClusterAddr, reqTimeout, int(poolSize))
		if redisClient == nil {
			panic("init redis cluster client failed.")
		} else {
			log.Infof("init redis cluster client : %s success.", redisClusterAddr)
		}
		svc.Global.RedisClient = redisClient

		rpcSvc := svc.NewRpcSvc()
		rpcSvc.Start(etcdAddrArr, tracerAddr)
	},
}
