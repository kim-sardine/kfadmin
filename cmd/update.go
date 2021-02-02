package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd update kubeflow resources
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update kubeflow resources",
	Long:  `TBU`,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
