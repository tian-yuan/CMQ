package commands

import (
	"strings"
	"time"

	"github.com/micro/go-micro/util/log"
	"github.com/spf13/cobra"
	"github.com/tian-yuan/CMQ/registry/svc"
	"github.com/tian-yuan/iot-common/util"
)

var rpccmd = &cobra.Command{
	Use:   "rpc",
	Short: "start rpc server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Start rpc server registry v0.0.1 -- HEAD")
		var zkAddr string
		cmd.Flags().StringVarP(&zkAddr, "zkAddress", "z", "127.0.0.1:2181", "zk address array")
		log.Infof("start register service, zk address : %s", zkAddr)
		zkAddrArr := strings.Split(zkAddr, ";")

		log.Info("begin to start redis client.")
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
		svc.Global.RedisClient = redisClient

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
		conf := svc.NewDeviceConfig()
		conf.Host = mysqlhost
		conf.Port = mysqlport
		conf.Username = username
		conf.Password = password
		conf.Database = database
		svc.Global.DeviceSvc = svc.NewDeviceSvc(conf)
		svc.Global.DeviceSvc.Start()

		rpcSvc := svc.NewRpcSvc()
		rpcSvc.Start(zkAddrArr)
	},
}
