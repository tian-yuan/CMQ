package commands

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"

	"github.com/tian-yuan/CMQ/topic-manager/svc"
	"time"
	"github.com/tian-yuan/iot-common/util"
)

var rpccmd = &cobra.Command{
	Use:   "rpc",
	Short: "start rpc server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start rpc server message dispatcher v0.0.1 -- HEAD")
		rpcSvc := svc.NewRpcSvc()
		rpcSvc.Start()

		var redisClusterAddr string
		cmd.Flags().StringVarP(&redisClusterAddr, "redisClusterAddr", "r", "127.0.0.1:7000", "redis cluster address")
		var reqTimeout time.Duration
		cmd.Flags().DurationVarP(&reqTimeout, "redisReqTimeout", "t", 5 * time.Second, "redis request timeout")
		var poolSize int32
		cmd.Flags().Int32VarP(&poolSize, "redisPoolSize", "s", 5, "redis client pool size")
		var redisSessionTimeout time.Duration
		cmd.Flags().DurationVarP(&redisSessionTimeout, "redisSessionTimeout", "o", 20 * time.Minute, "redis session timeout")
		var redisSessionRefresh time.Duration
		cmd.Flags().DurationVarP(&redisSessionRefresh, "redisSessionRefresh", "f", 20 * time.Minute, "redis session refresh")
		redisClient := util.GetClusterClient(redisClusterAddr, reqTimeout, int(poolSize))
		svc.Global.RedisClient = redisClient
	},
}
