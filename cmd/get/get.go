package get

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdGet(c *client.KfClient) *cobra.Command {

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get kubeflow resources",
		Long:  `TBU`,
	}

	getCmd.AddCommand(NewCmdGetProfiles(c))
	getCmd.AddCommand(NewCmdGetUsers(c))

	return getCmd
}
