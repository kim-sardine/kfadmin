package update

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdUpdateProfile(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Update kubeflow profile",
		Long:  `Update kubeflow profile`,
	}

	cmd.AddCommand(NewCmdUpdateProfileOwner(c))

	return cmd
}
