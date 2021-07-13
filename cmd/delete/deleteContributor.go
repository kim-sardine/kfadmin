package delete

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/kim-sardine/kfadmin/manifest"
)

type DeleteContributorOptions struct {
	Email   string
	Profile string

	clioption.IOStreams
}

func NewDeleteContributorOptions(ioStreams clioption.IOStreams) *DeleteContributorOptions {
	return &DeleteContributorOptions{
		IOStreams: ioStreams,
	}

}

func NewCmdDeleteContributor(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewDeleteContributorOptions(ioStreams)

	cmd := &cobra.Command{
		Use:   "contributor",
		Short: "Delete contributor from kubeflow profile",
		Long:  `Delete contributor from kubeflow profile`,
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

func (o *DeleteContributorOptions) Run(c *client.KfClient, cmd *cobra.Command) error {

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

	bindingName, err := manifest.GetBindingName(email)
	if err != nil {
		return err
	}

	err = c.DeleteRoleBinding(profile, bindingName)
	if err != nil {
		return err
	}
	err = c.DeleteServiceRoleBinding(profile, bindingName)
	if err != nil {
		return err
	}

	fmt.Fprintf(o.Out, "Contributor '%s' has been removed from '%s'\n", email, profile)

	return nil
}
