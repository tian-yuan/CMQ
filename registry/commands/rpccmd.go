package commands

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"

	"github.com/tian-yuan/CMQ/registry/svc"
	"strings"
)

var rpccmd = &cobra.Command{
	Use:   "rpc",
	Short: "start rpc server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start rpc server message dispatcher v0.0.1 -- HEAD")
		var zkAddr string
		cmd.Flags().StringVarP(&zkAddr, "zkAddress", "z", "127.0.0.1:2181", "zk address array")
		logrus.Infof("start register service, zk address : %s", zkAddr)
		zkAddrArr := strings.Split(zkAddr, ";")

		rpcSvc := svc.NewRpcSvc()
		rpcSvc.Start(zkAddrArr)
	},
}
