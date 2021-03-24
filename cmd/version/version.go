package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {

	var version string = "v21.03.25"

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	return cmd
}
