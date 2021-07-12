package create

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
)

// TODO: Add Option, Remove Dex

// NewCmdCreateContributor TBU
func NewCmdCreateContributor(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "contributor",
		Short: "Create contributor for kubeflow profile",
		Long:  `Create contributor for kubeflow profile`,
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
				panic(fmt.Errorf("User with email '%s' does not exist", email))
			}

			serviceRoleBindingManifest, err := manifest.GetServiceRoleBinding(profile, email)
			if err != nil {
				panic(err)
			}
			srb, err := c.GetServiceRoleBinding(profile, serviceRoleBindingManifest.Name)
			if err != nil {
				if !errors.IsNotFound(err) {
					panic(err)
				}
			}
			if srb != nil {
				panic(fmt.Errorf("ServiceRoleBinding '%s' already exists", serviceRoleBindingManifest.Name))
			}

			roleBinding, err := manifest.GetRoleBinding(profile, email)
			if err != nil {
				panic(err)
			}
			err = c.CreateRoleBinding(profile, roleBinding)
			if err != nil {
				panic(err)
			}
			err = c.CreateServiceRoleBinding(profile, serviceRoleBindingManifest)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Added '%s' to '%s'\n", email, profile)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringP("profile", "p", "", "Profile name")
	cmd.MarkFlagRequired("profile")
	cmd.Flags().StringP("email", "e", "", "User email")
	cmd.MarkFlagRequired("email")

	return cmd
}
