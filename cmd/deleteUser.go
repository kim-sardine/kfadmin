package cmd

import (
	"github.com/kim-sardine/kfadmin/admin"

	"github.com/spf13/cobra"
)

// deleteUserCmd delete new user
var deleteUserCmd = &cobra.Command{
	Use:   "user",
	Short: "delete new user",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")

		admin.DeleteUser(email)
	},
}

func init() {
	deleteCmd.AddCommand(deleteUserCmd)

	deleteUserCmd.Flags().SortFlags = false
	deleteUserCmd.Flags().StringP("email", "e", "", "User email")
	deleteUserCmd.MarkFlagRequired("email")
}
