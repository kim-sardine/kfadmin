package create

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
)

// createProfileCmd create kubeflow profile
var createProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "create kubeflow profile",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		profileName, _ := cmd.Flags().GetString("profile")
		email, _ := cmd.Flags().GetString("email")

		_, err := c.GetProfile(profileName)
		if err == nil {
			panic(fmt.Errorf("Profile '%s' already exists", profileName))
		}
		if !errors.IsNotFound(err) {
			panic(err)
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

		profile := manifest.GetProfile(profileName, email)
		err = c.CreateProfile(profile)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Profile '%s' created\n", profileName)
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
