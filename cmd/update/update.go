package update

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdUpdate(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "update",
		Short: "update kubeflow resources",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdUpdateProfile(c))
	cmd.AddCommand(NewCmdUpdateUser(c))

	return cmd
}
