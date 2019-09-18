package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init()  {
	rootCmd.AddCommand(cacheCmd)
}

var rootCmd = &cobra.Command{
	Use: "cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OK")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func Stop() {

}