package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd get kubeflow resources
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get kubeflow resources",
	Long:  `TBU`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
