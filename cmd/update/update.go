package update

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

func NewCmdUpdate(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update kubeflow resources",
		Long:  `Update kubeflow resources`,
	}

	cmd.AddCommand(NewCmdUpdateProfile(c))
	cmd.AddCommand(NewCmdUpdateStaticUser(c))

	return cmd
}
