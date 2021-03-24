package get

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdGet(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get kubeflow resources",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdGetProfiles(c))
	cmd.AddCommand(NewCmdGetUsers(c))

	return cmd
}
