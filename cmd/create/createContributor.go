package create

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/kim-sardine/kfadmin/manifest"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
)

// TODO: Add Option, Remove Dex
type CreateContributorOptions struct {
	Email   string
	Profile string

	clioption.IOStreams
}

func NewCreateContributorOptions(ioStreams clioption.IOStreams) *CreateContributorOptions {
	return &CreateContributorOptions{
		IOStreams: ioStreams,
	}

}

func NewCmdCreateContributor(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewCreateContributorOptions(ioStreams)

	cmd := &cobra.Command{
		Use:   "contributor",
		Short: "Create contributor for kubeflow profile",
		Long:  `Create contributor for kubeflow profile`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run(c, cmd))
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&o.Profile, "profile", "p", "", "Profile name")
	cmd.MarkFlagRequired("profile")
	cmd.Flags().StringVarP(&o.Email, "email", "e", "", "Email of new dex static user")
	cmd.MarkFlagRequired("email")

	return cmd
}

func (o *CreateContributorOptions) Run(c *client.KfClient, cmd *cobra.Command) error {

	profile, _ := cmd.Flags().GetString("profile")
	email, _ := cmd.Flags().GetString("email")

	_, err := c.GetProfile(profile)
	if err != nil {
		if errors.IsNotFound(err) {
			return fmt.Errorf("profile '%s' does not exists", profile)
		} else {
			return err
		}
	}

	users, err := c.GetStaticUsers()
	if err != nil {
		return err
	}
	userExists := false
	for _, user := range users {
		if user.Email == email {
			userExists = true
			break
		}
	}
	if !userExists {
		return fmt.Errorf("user with email '%s' does not exist", email)
	}

	serviceRoleBindingManifest, err := manifest.GetServiceRoleBinding(profile, email)
	if err != nil {
		return err
	}
	srb, err := c.GetServiceRoleBinding(profile, serviceRoleBindingManifest.Name)
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
	}
	if srb != nil {
		return fmt.Errorf("ServiceRoleBinding '%s' already exists", serviceRoleBindingManifest.Name)
	}

	roleBinding, err := manifest.GetRoleBinding(profile, email)
	if err != nil {
		return err
	}
	err = c.CreateRoleBinding(profile, roleBinding)
	if err != nil {
		return err
	}
	err = c.CreateServiceRoleBinding(profile, serviceRoleBindingManifest)
	if err != nil {
		return err
	}

	fmt.Fprintf(o.Out, "Added '%s' to '%s'\n", email, profile)

	return nil
}
