package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

type VersionString string

var CurrentVersion VersionString = "v21.3.26"

func NewCmdVersion() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(CurrentVersion)
		},
	}

	return cmd
}
