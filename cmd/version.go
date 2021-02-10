package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "v21.02.06"

// versionCmd print kfadmin version
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
