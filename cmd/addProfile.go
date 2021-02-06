package cmd

import (
	"github.com/spf13/cobra"
)

// addProfileCmd add resource to profile
var addProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "add resource to kubeflow profile",
	Long:  `TBU`,
}

func init() {
	addCmd.AddCommand(addProfileCmd)
}
