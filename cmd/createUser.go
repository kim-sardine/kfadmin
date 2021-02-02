package cmd

import (
	"github.com/kim-sardine/kfadmin/admin"

	"github.com/spf13/cobra"
)

// createUserCmd create new user
var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "create new user",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		admin.CreateUser(email, password)
	},
}

func init() {
	createCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().StringP("email", "e", "", "User email")
	createUserCmd.MarkFlagRequired("email")
	createUserCmd.Flags().StringP("password", "p", "", "User password")
	createUserCmd.MarkFlagRequired("password")
}
