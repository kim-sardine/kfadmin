package get

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

func NewCmdGet(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get kubeflow resources",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdGetProfiles(c))
	cmd.AddCommand(NewCmdGetDexUsers(c))

	return cmd
}
