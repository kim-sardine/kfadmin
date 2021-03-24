package delete

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdDelete(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete kubeflow resources",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdDeleteProfile(c))
	cmd.AddCommand(NewCmdDeleteUser(c))
	cmd.AddCommand(NewCmdDeleteContributor(c))

	return cmd
}
