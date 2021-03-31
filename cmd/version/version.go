package version

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

type VersionString string

var CurrentVersion VersionString = "v21.3.31"

func NewCmdVersion(ioStreams clioption.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(CurrentVersion)
		},
	}

	return cmd
}
