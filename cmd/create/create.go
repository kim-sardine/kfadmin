package create

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdCreate(c *client.KfClient) *cobra.Command {

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create kubeflow resources",
		Long:  `TBU`,
	}

	createCmd.AddCommand(NewCmdCreateProfile(c))
	createCmd.AddCommand(NewCmdCreateUser(c))

	return createCmd
}
