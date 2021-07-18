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

	clioption.DexOptions
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
	util.AddRestartDexFlag(cmd, &o.RestartDex)

	return cmd
}

func (o *UpdateStaticUserPasswordOptions) Run(c *client.KfClient, cmd *cobra.Command) error {

	email, _ := cmd.Flags().GetString("email")
	password, _ := cmd.Flags().GetString("password")

	// Check if user exists
	cm, err := c.GetDexConfigMap()
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

	if err := c.UpdateDexConfigMap(cm); err != nil {
		return err
	}

	fmt.Fprintf(o.Out, "password of user '%s' has been updated\n", email)

	if o.RestartDex {
		if err := c.RestartDex(o.ErrOut, originalData); err != nil {
			return err
		}
	} else {
		fmt.Fprintf(o.Out, "To reflect changes, please run below command\n\nkubectl rollout restart deployment dex -n auth\n\n")
	}

	return nil
}
