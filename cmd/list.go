package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd list kubeflow resources
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list kubeflow resources",
	Long:  `TBU`,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
