package update

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

func NewCmdUpdateProfile(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Update kubeflow profile",
		Long:  `Update kubeflow profile`,
	}

	cmd.AddCommand(NewCmdUpdateProfileOwner(c, ioStreams))

	return cmd
}
