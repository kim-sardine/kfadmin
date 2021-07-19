package update

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/kim-sardine/kfadmin/manifest"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
)

type UpdateProfileOwnerOptions struct {
	Profile string
	Email   string

	clioption.IOStreams
}

func NewUpdateProfileOwnerOptions(ioStreams clioption.IOStreams) *UpdateProfileOwnerOptions {
	return &UpdateProfileOwnerOptions{
		IOStreams: ioStreams,
	}
}

func NewCmdUpdateProfileOwner(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewUpdateProfileOwnerOptions(ioStreams)

	cmd := &cobra.Command{
		Use:   "owner",
		Short: "Update kubeflow profile owner",
		Long:  `Update kubeflow profile owner`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run(c, cmd))
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&o.Profile, "profile", "p", "", "Profile name")
	cmd.MarkFlagRequired("profile")
	cmd.Flags().StringVarP(&o.Email, "email", "e", "", "Owner's email")
	cmd.MarkFlagRequired("email")

	return cmd
}

func (o *UpdateProfileOwnerOptions) Run(c *client.KfClient, cmd *cobra.Command) error {
	// Case 1. contributor to owner -> need to delete contributor-rolebinding
	// Case 2. non-contributor to owner -> just change existing owner-rolebinding

	profileName, _ := cmd.Flags().GetString("profile")
	email, _ := cmd.Flags().GetString("email")

	profile, err := c.GetProfile(profileName)
	if err != nil {
		if errors.IsNotFound(err) {
			return fmt.Errorf("profile '%s' does not exists", profileName)
		} else {
			return err
		}
	}

	profile.Spec.Owner.Name = email
	if err := c.UpdateProfile(profile); err != nil {
		return err
	}

	ownerRoleBinding, err := c.GetRoleBinding(profileName, "namespaceAdmin")
	if err != nil {
		return err
	}
	ownerRoleBinding.Annotations["user"] = email
	ownerRoleBinding.Subjects[0].Name = email
	if err := c.UpdateRoleBinding(profileName, ownerRoleBinding); err != nil {
		return err
	}

	// FIXME: Not working here in kubeflow v1.3. Check Authorizationpolicy
	ownerServiceRoleBinding, err := c.GetServiceRoleBinding(profileName, "owner-binding-istio")
	if err != nil {
		return err
	}
	ownerServiceRoleBinding.Annotations["user"] = email
	ownerServiceRoleBinding.Spec.Subjects[0].Properties["request.headers[kubeflow-userid]"] = email
	if err := c.UpdateServiceRoleBinding(profileName, ownerServiceRoleBinding); err != nil {
		return err
	}

	// Delete old contributor-roleBinding
	bindingName, err := manifest.GetBindingName(email)
	if err != nil {
		return err
	}

	contributorRoleBinding, err := c.GetRoleBinding(profileName, bindingName)
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
	}

	if contributorRoleBinding != nil { // Case 1
		err = c.DeleteRoleBinding(profileName, bindingName)
		if err != nil {
			return err
		}
		err = c.DeleteServiceRoleBinding(profileName, bindingName)
		if err != nil {
			return err
		}
	}

	fmt.Fprintf(o.Out, "Owner of the profile '%s' has changed to '%s'\n", profileName, email)

	return nil
}
