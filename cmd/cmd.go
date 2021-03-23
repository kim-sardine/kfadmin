package cmd

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/cmd/create"
	"github.com/spf13/cobra"
)

var cfgFile string

var c *client.KfClient

func NewKfAdminCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "kfadmin",
		Short: "CLI Tool for Kubeflow admin",
		Long: `kfadmin is a CLI tool for kubeflow admin.

Find more information at: https://github.com/kim-sardine/kfadmin

Examples:
- kfadmin create user -e USER_EMAIL -p PASSWORD
- kfadmin get user
- kfadmin delete user -e USER_EMAIL
- kfadmin update user password -e USER_EMAIL -p NEW_PASSWORD
- kfadmin create profile -p PROFILE_NAME -e OWNER_EMAIL
- kfadmin get profile
- kfadmin delete namespace -p PROFILE_NAME
- kfadmin add profile contributor -p PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL
- kfadmin update profile owner -p PROFILE_NAME -e NEW_OWNER_EMAIL
- kfadmin delete profile contributor -p PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL`,
	}

	kfClient := &client.KfClient{}
	kfClient.LoadClientset()

	rootCmd.AddCommand(create.NewCmdCreate(kfClient))

	return rootCmd
}
