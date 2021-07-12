package cmd

import (
	"io"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/completion"
	"github.com/kim-sardine/kfadmin/cmd/create"
	"github.com/kim-sardine/kfadmin/cmd/delete"
	"github.com/kim-sardine/kfadmin/cmd/get"
	"github.com/kim-sardine/kfadmin/cmd/update"
	"github.com/kim-sardine/kfadmin/cmd/version"
	"github.com/spf13/cobra"
)

func NewKfAdminCommand(in io.Reader, out, err io.Writer) *cobra.Command {

	var rootCmd = &cobra.Command{
		Use:   "kfadmin",
		Short: "CLI Tool for Kubeflow admin",
		Long: `kfadmin is a CLI tool for kubeflow admin.

Find more information at: https://github.com/kim-sardine/kfadmin

Examples:

- kfadmin get staticusers
- kfadmin get profiles

- kfadmin create staticuser -e USER_EMAIL -p PASSWORD
- kfadmin create profile -p PROFILE_NAME -e OWNER_EMAIL
- kfadmin create contributor -p PROFILE_NAME -e USER_EMAIL

- kfadmin update staticuser password -e USER_EMAIL -p NEW_PASSWORD
- kfadmin update profile owner -p PROFILE_NAME -e NEW_OWNER_EMAIL

- kfadmin delete staticuser -e USER_EMAIL
- kfadmin delete profile -p PROFILE_NAME
- kfadmin delete contributor -p PROFILE_NAME -e NEW_CONTRIBUTOR_EMAIL`,
	}

	ioStreams := clioption.IOStreams{In: in, Out: out, ErrOut: err}

	kfClient := client.NewKfClient()
	kfClient.LoadClientset()

	rootCmd.AddCommand(create.NewCmdCreate(kfClient, ioStreams))
	rootCmd.AddCommand(get.NewCmdGet(kfClient, ioStreams))
	rootCmd.AddCommand(update.NewCmdUpdate(kfClient, ioStreams))
	rootCmd.AddCommand(delete.NewCmdDelete(kfClient, ioStreams))

	rootCmd.AddCommand(completion.NewCmdCompletion(ioStreams.Out))
	rootCmd.AddCommand(version.NewCmdVersion(ioStreams))

	return rootCmd
}
