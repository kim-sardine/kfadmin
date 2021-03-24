package cmd

import (
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/cmd/create"
	"github.com/kim-sardine/kfadmin/cmd/get"
	"github.com/spf13/cobra"
)

func NewKfAdminCommand() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:   "kfadmin",
		Short: "CLI Tool for Kubeflow admin",
		Long: `kfadmin is a CLI tool for kubeflow admin.

Find more information at: https://github.com/kim-sardine/kfadmin

Examples:

- kfadmin get users
- kfadmin get profiles

- kfadmin create user -e USER_EMAIL -p PASSWORD
- kfadmin create profile -p PROFILE_NAME -e OWNER_EMAIL

- kfadmin add profile contributor -p PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL

- kfadmin update user password -e USER_EMAIL -p NEW_PASSWORD
- kfadmin update profile owner -p PROFILE_NAME -e NEW_OWNER_EMAIL

- kfadmin delete user -e USER_EMAIL
- kfadmin delete namespace -p PROFILE_NAME
- kfadmin delete profile contributor -p PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL`,
	}

	kfClient := client.NewKfClient()
	kfClient.LoadClientset()

	rootCmd.AddCommand(create.NewCmdCreate(kfClient))
	rootCmd.AddCommand(get.NewCmdGet(kfClient))

	return rootCmd
}
