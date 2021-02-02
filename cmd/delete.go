package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd delete kubeflow resources
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete kubeflow resources",
	Long:  `TBU`,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
