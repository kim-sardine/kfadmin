package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd create kubeflow resources
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create kubeflow resources",
	Long:  `TBU`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
