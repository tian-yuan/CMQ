package commands

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"

	"github.com/tian-yuan/CMQ/publish-engine/svc"
)

var httpCmd = &cobra.Command{
	Use:   "rpc",
	Short: "start rpc server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start rpc server message dispatcher v0.0.1 -- HEAD")
		rpcSvc := svc.NewRpcSvc()
		rpcSvc.Start()
	},
}
