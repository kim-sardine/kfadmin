package update

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
)

func NewCmdUpdateProfileOwner(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "owner",
		Short: "update kubeflow profile owner",
		Long:  `TBU`,
		Run: func(cmd *cobra.Command, args []string) {
			profileName, _ := cmd.Flags().GetString("profile")
			email, _ := cmd.Flags().GetString("email")

			profile, err := c.GetProfile(profileName)
			if err != nil {
				if errors.IsNotFound(err) {
					panic(fmt.Errorf("Profile '%s' does not exists", profileName))
				} else {
					panic(err)
				}
			}

			// Update existing resources
			// profile
			profile.Spec.Owner.Name = email
			err = c.UpdateProfile(profile)
			if err != nil {
				panic(err)
			}

			// rbacv1.RoleBinding namespaceAdmin
			rb, err := c.GetRoleBinding(profileName, "namespaceAdmin")
			if err != nil {
				panic(err)
			}

			rb.Annotations["user"] = email
			rb.Subjects[0].Name = email
			err = c.UpdateRoleBinding(profileName, rb)
			if err != nil {
				panic(err)
			}

			// istiorbac.ServiceRoleBinding
			srb, err := c.GetServiceRoleBinding(profileName, "owner-binding-istio")
			if err != nil {
				panic(err)
			}

			srb.Annotations["user"] = email
			srb.Spec.Subjects[0].Properties["request.headers[kubeflow-userid]"] = email

			err = c.UpdateServiceRoleBinding(profileName, srb)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Owner of the profile '%s' has changed to '%s'\n", profileName, email)

		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringP("profile", "p", "", "Profile name")
	cmd.MarkFlagRequired("profile")
	cmd.Flags().StringP("email", "e", "", "New owner's email")
	cmd.MarkFlagRequired("email")

	return cmd
}
