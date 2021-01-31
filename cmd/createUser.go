package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kim-sardine/kfadmin/admin"
)

// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		admin.CreateUser(email, password)
	},
}

func init() {
	rootCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().StringP("email", "e", "", "User email")
	createUserCmd.MarkFlagRequired("email")
	createUserCmd.Flags().StringP("password", "p", "", "User password")
	createUserCmd.MarkFlagRequired("password")
}
