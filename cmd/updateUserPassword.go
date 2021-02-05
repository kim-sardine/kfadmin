package cmd

import (
	"github.com/kim-sardine/kfadmin/admin"

	"github.com/spf13/cobra"
)

// updateUserPasswordCmd udpate staticPassword's password
var updateUserPasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "udpate dex staticPassword's password",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")
		admin.UpdateUserPassword(email, password)
	},
}

func init() {
	updateUserCmd.AddCommand(updateUserPasswordCmd)

	updateUserPasswordCmd.Flags().SortFlags = false
	updateUserPasswordCmd.Flags().StringP("email", "e", "", "User email")
	updateUserPasswordCmd.MarkFlagRequired("email")
	updateUserPasswordCmd.Flags().StringP("password", "p", "", "User password")
	updateUserPasswordCmd.MarkFlagRequired("password")
}
