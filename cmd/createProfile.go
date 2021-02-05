package cmd

import (
	"github.com/kim-sardine/kfadmin/admin"

	"github.com/spf13/cobra"
)

// createProfileCmd create kubeflow profile
var createProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "create kubeflow profile",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		profile, _ := cmd.Flags().GetString("profile")
		email, _ := cmd.Flags().GetString("email")

		admin.CreateProfile(profile, email)
	},
}

func init() {
	createCmd.AddCommand(createProfileCmd)

	createProfileCmd.Flags().SortFlags = false
	createProfileCmd.Flags().StringP("profile", "p", "", "Profile name")
	createProfileCmd.MarkFlagRequired("profile")
	createProfileCmd.Flags().StringP("email", "e", "", "Owner email")
	createProfileCmd.MarkFlagRequired("email")

}
