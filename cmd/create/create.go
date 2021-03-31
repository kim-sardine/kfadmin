package create

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

func NewCmdCreate(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create kubeflow resources",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdCreateProfile(c))
	cmd.AddCommand(NewCmdCreateUser(c))
	cmd.AddCommand(NewCmdCreateContributor(c))

	return cmd
}
