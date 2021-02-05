package cmd

import (
	"github.com/kim-sardine/kfadmin/admin"

	"github.com/spf13/cobra"
)

// listProfileCmd create kubeflow profile
var listProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "print kubeflow profile",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		admin.ListProfile()
	},
}

func init() {
	listCmd.AddCommand(listProfileCmd)
}
