package update

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdUpdateDexUser(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "dex-user",
		Short: "update user information",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdUpdateDexUserPassword(c))

	return cmd
}
