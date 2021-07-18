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

type CreateProfileOptions struct {
	Profile string
	Email   string

	clioption.IOStreams
}

func NewCreateProfileOptions(ioStreams clioption.IOStreams) *CreateProfileOptions {
	return &CreateProfileOptions{
		IOStreams: ioStreams,
	}
}

func NewCmdCreateProfile(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewCreateProfileOptions(ioStreams)

	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Create kubeflow profile",
		Long:  `Create kubeflow profile`,
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

func (o *CreateProfileOptions) Run(c *client.KfClient, cmd *cobra.Command) error {

	profileName, _ := cmd.Flags().GetString("profile")
	email, _ := cmd.Flags().GetString("email")

	_, err := c.GetProfile(profileName)
	if err == nil {
		return fmt.Errorf("profile '%s' already exists", profileName)
	}
	if !errors.IsNotFound(err) {
		return err
	}

	if _, err = c.GetDexConfigMap(); err == nil {
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
			return fmt.Errorf("user : '%s' does not exist", email)
		}
	}

	profile := manifest.GetProfile(profileName, email)

	if err := c.CreateProfile(profile); err != nil {
		return err
	}

	fmt.Fprintf(o.Out, "Profile '%s' created\n", profileName)

	return nil
}
