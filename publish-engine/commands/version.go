package commands

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Hub gateway v0.0.1 -- HEAD")
	},
}
