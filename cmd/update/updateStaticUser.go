package update

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdUpdateStaticUser(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "static-user",
		Short: "Update dex static user information",
		Long:  `Update dex static user information`,
	}

	cmd.AddCommand(NewCmdUpdateStaticUserPassword(c))

	return cmd
}
