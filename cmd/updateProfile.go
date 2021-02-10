package cmd

import (
	"github.com/spf13/cobra"
)

// updateProfileCmd update kubeflow profile
var updateProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "update kubeflow profile",
	Long:  `TBU`,
}

func init() {
	updateCmd.AddCommand(updateProfileCmd)
}
