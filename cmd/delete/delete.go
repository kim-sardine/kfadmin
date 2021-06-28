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
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdDeleteProfile(c))
	cmd.AddCommand(NewCmdDeleteStaticUser(c))
	cmd.AddCommand(NewCmdDeleteContributor(c))

	return cmd
}
