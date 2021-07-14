package create

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/kim-sardine/kfadmin/manifest"
	"github.com/spf13/cobra"
)

// TODO: Move RestartDex option to shared dex option
type CreateStaticUserOptions struct {
	Email    string
	Password string

	RestartDex bool

	clioption.IOStreams
}

// NewCreateStaticUserOptions returns initialized Options
func NewCreateStaticUserOptions(ioStreams clioption.IOStreams) *CreateStaticUserOptions {
	return &CreateStaticUserOptions{
		IOStreams: ioStreams,
	}

}

func NewCmdCreateStaticUser(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewCreateStaticUserOptions(ioStreams)

	cmd := &cobra.Command{
		Use:     "staticuser",
		Aliases: []string{"static-user"},
		Short:   "Create new dex static user",
		Long:    `Create new dex static user`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run(c, cmd))
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&o.Email, "email", "e", "", "Email of new dex static user")
	cmd.MarkFlagRequired("email")
	cmd.Flags().StringVarP(&o.Password, "password", "p", "", "Password of new dex static user")
	cmd.MarkFlagRequired("password")
	// TODO: into shared util
	cmd.Flags().BoolVarP(&o.RestartDex, "restart-dex", "r", false, "Restart dex deployment to reflect changes")

	return cmd
}

func (o *CreateStaticUserOptions) Run(c *client.KfClient, cmd *cobra.Command) error {

	email := o.Email
	password := o.Password
	restartDex := o.RestartDex

	username, err := util.GetUsernameFromEmail(email)
	if err != nil {
		return err
	}

	cm, err := c.GetDexConfigMap()
	if err != nil {
		return err
	}

	originalData := cm.Data["config.yaml"]
	dexDataConfig, err := manifest.UnmarshalDexDataConfig(originalData)
	if err != nil {
		return err
	}

	users := dexDataConfig.StaticPasswords

	uuids := make([]string, len(users)+1)
	for _, user := range users {
		uuids = append(uuids, user.UserID)
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	newUser := manifest.StaticPassword{
		Email:    email,
		Hash:     hashedPassword,
		Username: username,
		UserID:   util.GetUniqueUUID(uuids),
	}

	dexDataConfig.StaticPasswords = append(dexDataConfig.StaticPasswords, newUser)

	cm.Data["config.yaml"], err = manifest.MarshalDexDataConfig(dexDataConfig)
	if err != nil {
		return err
	}

	err = c.UpdateDexConfigMap(cm)
	if err != nil {
		return err
	}

	// TODO: into shared util
	if restartDex {
		if err = c.RestartDexDeployment(); err != nil {
			fmt.Fprintf(o.ErrOut, err.Error()+"\n")
			fmt.Fprintf(o.ErrOut, "failed to restart dex deployment, rollback dex deployment...\n")
			if err = c.RollbackDexDeployment(originalData); err != nil {
				return err
			}
			return fmt.Errorf("completed rollback dex deployment")
		}

		fmt.Fprintf(o.Out, "user '%s' created\n", email)
	} else {
		fmt.Fprintf(o.Out, "user '%s' has been added to dex\nTo reflect changes, please run below command\n\nkubectl rollout restart deployment dex -n auth\n\n", email)
	}

	return nil
}
