package cmd

import (
	"github.com/kim-sardine/kfadmin/admin"

	"github.com/spf13/cobra"
)

// listUserCmd list all user
var listUserCmd = &cobra.Command{
	Use:   "user",
	Short: "list all user",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {

		admin.ListUser()
	},
}

func init() {
	listCmd.AddCommand(listUserCmd)
}
