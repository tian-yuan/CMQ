package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tian-yuan/CMQ/controller/svc"
)

var httpCmd = &cobra.Command{
	Use:   "controller",
	Short: "start controller server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start controller server v0.0.1 -- HEAD")
		httpconf := svc.NewH2cConf()
		cmd.Flags().StringVarP(&httpconf.Host, "Host", "h", "0.0.0.0", "http2 bind host address.")
		cmd.Flags().Uint16VarP(&httpconf.Port, "Port", "p", 9883, "http2 hub bind port.")
		h2cSvc := svc.NewH2cSvc(httpconf)
		h2cSvc.Start()

		dbconf := svc.NewDatabaseConfig()
		cmd.Flags().StringVarP(&dbconf.Host, "Host", "m", "0.0.0.0", "mysql host address")
		cmd.Flags().Uint16VarP(&dbconf.Port, "Port", "p", 3306, "mysql port")
		cmd.Flags().StringVarP(&dbconf.Username, "Username", "u", "root", "mysql username")
		cmd.Flags().StringVarP(&dbconf.Password, "password", "w", "123456", "mysql password")
		cmd.Flags().StringVarP(&dbconf.Database, "database", "d", "device", "mysql database")
		dbsvc := svc.NewDatabaseSvc(dbconf)
		svc.Ctx.Dbsvc = dbsvc
		dbsvc.Start()
	},
}
