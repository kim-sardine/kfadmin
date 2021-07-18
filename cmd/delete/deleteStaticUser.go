package delete

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/kim-sardine/kfadmin/manifest"
	"github.com/spf13/cobra"
)

type DeleteStaticUserOptions struct {
	Email string

	clioption.DexOptions
	clioption.IOStreams
}

func NewDeleteStaticUserOptions(ioStreams clioption.IOStreams) *DeleteStaticUserOptions {
	return &DeleteStaticUserOptions{
		IOStreams: ioStreams,
	}

}

func NewCmdDeleteStaticUser(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewDeleteStaticUserOptions(ioStreams)

	cmd := &cobra.Command{
		Use:     "staticuser",
		Aliases: []string{"static-user"},
		Short:   "Delete dex static user",
		Long:    `Delete dex static user`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run(c, cmd))
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&o.Email, "email", "e", "", "Email of new dex static user")
	cmd.MarkFlagRequired("email")
	util.AddRestartDexFlag(cmd, &o.RestartDex)

	return cmd
}

func (o *DeleteStaticUserOptions) Run(c *client.KfClient, cmd *cobra.Command) error {

	email, _ := cmd.Flags().GetString("email")

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

	// Delete if exists
	dc.StaticPasswords = append(dc.StaticPasswords[:userIdx], dc.StaticPasswords[userIdx+1:]...)
	cm.Data["config.yaml"], err = manifest.MarshalDexDataConfig(dc)
	if err != nil {
		return err
	}

	if err := c.UpdateDexConfigMap(cm); err != nil {
		return err
	}

	fmt.Fprintf(o.Out, "static user '%s' has been deleted successfully\n", email)

	if o.RestartDex {
		if err := c.RestartDex(o.ErrOut, originalData); err != nil {
			return err
		}
	} else {
		fmt.Fprintf(o.Out, "To reflect changes, please run below command\n\nkubectl rollout restart deployment dex -n auth\n\n")
	}

	return nil
}
