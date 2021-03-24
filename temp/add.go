package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd add kubeflow resources
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add kubeflow resources",
	Long:  `TBU`,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
