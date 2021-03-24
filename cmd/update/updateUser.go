package update

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdUpdateUser(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "update user information",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdUpdateUserPassword(c))

	return cmd
}
