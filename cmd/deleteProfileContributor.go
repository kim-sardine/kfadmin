package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/kim-sardine/kfadmin/client/manifest"
)

// deleteProfileContributorCmd add contributor to kubeflow profile
var deleteProfileContributorCmd = &cobra.Command{
	Use:   "contributor",
	Short: "delete contributor from kubeflow profile",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		profile, _ := cmd.Flags().GetString("profile")
		email, _ := cmd.Flags().GetString("email")

		_, err := c.GetProfile(profile)
		if err != nil {
			if errors.IsNotFound(err) {
				panic(fmt.Errorf("Profile '%s' does not exists", profile))
			} else {
				panic(err)
			}
		}

		users, err := c.GetStaticUsers()
		if err != nil {
			panic(err)
		}
		userExists := false
		for _, user := range users {
			if user.Email == email {
				userExists = true
				break
			}
		}
		if !userExists {
			panic(fmt.Errorf("User with email '%s' does not exist", email))
		}

		bindingName, err := manifest.GetBindingName(email)
		if err != nil {
			panic(err)
		}

		err = c.DeleteRoleBinding(profile, bindingName)
		if err != nil {
			panic(err)
		}
		err = c.DeleteServiceRoleBinding(profile, bindingName)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Contributor '%s' removed from '%s'\n", email, profile)
	},
}

func init() {
	deleteProfileCmd.AddCommand(deleteProfileContributorCmd)

	deleteProfileContributorCmd.Flags().SortFlags = false
	deleteProfileContributorCmd.Flags().StringP("profile", "p", "", "Profile name")
	deleteProfileContributorCmd.MarkFlagRequired("profile")
	deleteProfileContributorCmd.Flags().StringP("email", "e", "", "Owner email")
	deleteProfileContributorCmd.MarkFlagRequired("email")

}
