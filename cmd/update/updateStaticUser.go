package update

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdUpdateStaticUser(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "static-user",
		Short: "update dex static user information",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdUpdateStaticUserPassword(c))

	return cmd
}
