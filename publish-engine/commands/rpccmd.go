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
		var etcdAddr string
		cmd.Flags().StringVarP(&etcdAddr, "etcdAddress", "z", "127.0.0.1:2379", "etcd address array")
		log.Infof("start publish engine service, etcd address : %s", etcdAddr)
		etcdAddrArr := strings.Split(etcdAddr, ";")
		var tracerAddr string
		cmd.Flags().StringVarP(&tracerAddr, "tracerAddress", "j", "127.0.0.1:6831", "tracer address array")

		util.Init(
			util.WithRegistryUrls(etcdAddr),
			util.WithTracerUrl(tracerAddr),
		)
		util.Ctx.InitMessageDispatcherSvc()
		// close the client tracer
		defer util.Ctx.CloseMessageDispatcherSvc()

		rpcSvc := svc.NewRpcSvc()
		rpcSvc.Start(etcdAddrArr, tracerAddr)
	},
}
