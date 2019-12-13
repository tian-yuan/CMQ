package commands

import (
	"github.com/spf13/cobra"

	"strings"

	"github.com/micro/go-micro/util/log"
	"github.com/tian-yuan/CMQ/publish-engine/svc"
	"github.com/tian-yuan/iot-common/util"
)

var httpCmd = &cobra.Command{
	Use:   "rpc",
	Short: "start rpc server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Start rpc server message dispatcher v0.0.1 -- HEAD")
		var zkAddr string
		cmd.Flags().StringVarP(&zkAddr, "zkAddress", "z", "127.0.0.1:2181", "zk address array")
		log.Infof("start publish engine service, zk address : %s", zkAddr)
		zkAddrArr := strings.Split(zkAddr, ";")
		util.Ctx.InitMessageDispatcherSvc(zkAddrArr)

		rpcSvc := svc.NewRpcSvc()
		rpcSvc.Start(zkAddrArr)
	},
}
