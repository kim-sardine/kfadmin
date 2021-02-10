package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateProfileOwnerCmd update profile owner
var updateProfileOwnerCmd = &cobra.Command{
	Use:   "owner",
	Short: "update kubeflow profile owner",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wow")
	},
}

func init() {
	updateProfileCmd.AddCommand(updateProfileOwnerCmd)
}
