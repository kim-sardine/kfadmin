package update

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/kim-sardine/kfadmin/manifest"
	"github.com/spf13/cobra"
)

type UpdateStaticUserPasswordOptions struct {
	Email    string
	Password string

	clioption.IOStreams
}

func NewUpdateStaticUserPasswordOptions(ioStreams clioption.IOStreams) *UpdateStaticUserPasswordOptions {
	return &UpdateStaticUserPasswordOptions{
		IOStreams: ioStreams,
	}
}

func NewCmdUpdateStaticUserPassword(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewUpdateStaticUserPasswordOptions(ioStreams)

	cmd := &cobra.Command{
		Use:   "password",
		Short: "Udpate dex static user's password",
		Long:  `Udpate dex static user's password`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run(c, cmd))
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&o.Email, "email", "e", "", "Owner's email")
	cmd.MarkFlagRequired("email")
	cmd.Flags().StringVarP(&o.Password, "password", "p", "", "User password")
	cmd.MarkFlagRequired("password")

	return cmd
}

func (o *UpdateStaticUserPasswordOptions) Run(c *client.KfClient, cmd *cobra.Command) error {

	email, _ := cmd.Flags().GetString("email")
	password, _ := cmd.Flags().GetString("password")

	// Check if user exists
	cm, err := c.GetDex()
	if err != nil {
		return err
	}

	originalData := cm.Data["config.yaml"]
	dc, err := manifest.UnmarshalDexDataConfig(originalData)
	if err != nil {
		return err
	}

	users := dc.StaticPasswords
	userIdx := -1
	for i, user := range users {
		if user.Email == email {
			userIdx = i
			break
		}
	}
	if userIdx < 0 {
		return fmt.Errorf("user with email '%s' does not exist", email)
	}

	// Change Password
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	dc.StaticPasswords[userIdx].Hash = hashedPassword
	cm.Data["config.yaml"], err = manifest.MarshalDexDataConfig(dc)
	if err != nil {
		return err
	}

	err = c.UpdateDex(cm)
	if err != nil {
		return err
	}

	err = c.RestartDexDeployment(originalData)
	if err != nil {
		return err
	}

	fmt.Fprintf(o.Out, "user '%s' updated\n", email)

	return nil
}
