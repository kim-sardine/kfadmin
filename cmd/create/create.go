package create

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdCreate(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create kubeflow resources",
		Long:  `TBU`,
	}

	cmd.AddCommand(NewCmdCreateProfile(c))
	cmd.AddCommand(NewCmdCreateUser(c))
	cmd.AddCommand(NewCmdCreateContributor(c))

	return cmd
}
