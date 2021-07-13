package delete

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

func NewCmdDelete(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete kubeflow resources",
		Long:  `Delete kubeflow resources`,
	}

	cmd.AddCommand(NewCmdDeleteProfile(c, ioStreams))
	cmd.AddCommand(NewCmdDeleteStaticUser(c, ioStreams))
	cmd.AddCommand(NewCmdDeleteContributor(c, ioStreams))

	return cmd
}
