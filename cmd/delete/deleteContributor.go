package delete

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/manifest"
)

func NewCmdDeleteContributor(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "contributor",
		Short: "Delete contributor from kubeflow profile",
		Long:  `Delete contributor from kubeflow profile`,
		Run: func(cmd *cobra.Command, args []string) {
			profile, _ := cmd.Flags().GetString("profile")
			email, _ := cmd.Flags().GetString("email")

			_, err := c.GetProfile(profile)
			if err != nil {
				if errors.IsNotFound(err) {
					panic(fmt.Errorf("profile '%s' does not exists", profile))
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
				panic(fmt.Errorf("user with email '%s' does not exist", email))
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

	cmd.Flags().SortFlags = false
	cmd.Flags().StringP("profile", "p", "", "Profile name")
	cmd.MarkFlagRequired("profile")
	cmd.Flags().StringP("email", "e", "", "Owner email")
	cmd.MarkFlagRequired("email")

	return cmd
}
