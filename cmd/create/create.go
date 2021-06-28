package create

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

func NewCmdCreate(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create resources like kubeflow profile, dex static user. Or Add contributor",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdCreateProfile(c, ioStreams))
	cmd.AddCommand(NewCmdCreateStaticUser(c, ioStreams))
	cmd.AddCommand(NewCmdCreateContributor(c, ioStreams))

	return cmd
}
