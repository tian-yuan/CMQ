package commands

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "start cache task",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start cache task -- HEAD")
	},
}
