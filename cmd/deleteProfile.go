package cmd

import (
	"github.com/kim-sardine/kfadmin/admin"

	"github.com/spf13/cobra"
)

// deleteProfileCmd delete kubeflow profile
var deleteProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "delete kubeflow profile",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		profile, _ := cmd.Flags().GetString("profile")

		admin.DeleteProfile(profile)
	},
}

func init() {
	deleteCmd.AddCommand(deleteProfileCmd)

	deleteProfileCmd.Flags().SortFlags = false
	deleteProfileCmd.Flags().StringP("profile", "p", "", "Profile name")
	deleteProfileCmd.MarkFlagRequired("profile")

}
