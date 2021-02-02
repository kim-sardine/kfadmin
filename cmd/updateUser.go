package cmd

import (
	"github.com/kim-sardine/kfadmin/admin"

	"github.com/spf13/cobra"
)

// updateUserCmd set all user
var updateUserCmd = &cobra.Command{
	Use:   "user",
	Short: "set all user",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")
		admin.UpdateUser(email, password)
	},
}

func init() {
	updateCmd.AddCommand(updateUserCmd)

	updateUserCmd.Flags().StringP("email", "e", "", "User email")
	updateUserCmd.MarkFlagRequired("email")
	updateUserCmd.Flags().StringP("password", "p", "", "User password")
	updateUserCmd.MarkFlagRequired("password")
}
