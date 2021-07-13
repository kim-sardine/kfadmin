package update

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

func NewCmdUpdateStaticUser(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:     "staticuser",
		Aliases: []string{"static-user"},
		Short:   "Update dex static user information",
		Long:    `Update dex static user information`,
	}

	cmd.AddCommand(NewCmdUpdateStaticUserPassword(c))

	return cmd
}
