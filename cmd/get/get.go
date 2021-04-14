package get

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

func NewCmdGet(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Print the information about kubeflow resources",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdGetProfiles(c, ioStreams))
	cmd.AddCommand(NewCmdGetDexUsers(c, ioStreams))

	return cmd
}
