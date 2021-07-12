package create

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
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
		Use:   "static-user",
		Short: "Create new dex static user",
		Long:  `Create new dex static user`,
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

	dc.StaticPasswords = append(dc.StaticPasswords, newUser)

	cm.Data["config.yaml"], err = manifest.MarshalDexDataConfig(dc)
	if err != nil {
		return err
	}

	err = c.UpdateDex(cm)
	if err != nil {
		return err
	}

	// TODO: into shared util
	if restartDex {
		err = c.RestartDexDeployment(originalData)
		if err != nil {
			return err
		}
		fmt.Fprintf(o.Out, "user '%s' created\n", email)
	} else {
		fmt.Fprintf(o.Out, "user '%s' added to dex ConfigMap\nTo reflext changes, please run below command\n\nkubectl rollout restart deployment dex -n auth\n\n", email)
	}

	return nil
}
