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
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdUpdateProfile(c))
	cmd.AddCommand(NewCmdUpdateDexUser(c))

	return cmd
}
