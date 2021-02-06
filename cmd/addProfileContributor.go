package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/kim-sardine/kfadmin/client/manifest"
)

// addProfileContributorCmd add contributor to kubeflow profile
var addProfileContributorCmd = &cobra.Command{
	Use:   "contributor",
	Short: "add contributor to kubeflow profile",
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

		// TODO: Check below resource exists
		// TODO: Atomic Creation
		roleBinding, err := manifest.GetRoleBinding(profile, email)
		err = c.CreateRoleBinding(profile, roleBinding)
		if err != nil {
			panic(err)
		}
		serviceRoleBinding, err := manifest.GetServiceRoleBinding(profile, email)
		err = c.CreateServiceRoleBinding(profile, serviceRoleBinding)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Added '%s' to '%s'\n", email, profile)
	},
}

func init() {
	addProfileCmd.AddCommand(addProfileContributorCmd)

	addProfileContributorCmd.Flags().SortFlags = false
	addProfileContributorCmd.Flags().StringP("profile", "p", "", "Profile name")
	addProfileContributorCmd.MarkFlagRequired("profile")
	addProfileContributorCmd.Flags().StringP("email", "e", "", "Owner email")
	addProfileContributorCmd.MarkFlagRequired("email")

}
