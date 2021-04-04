package version

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

type VersionString string

var CurrentVersion VersionString = "v21.3.31"

type Options struct {
	clioption.IOStreams
}

// NewOptions returns initialized Options
func NewOptions(ioStreams clioption.IOStreams) *Options {
	return &Options{
		IOStreams: ioStreams,
	}

}

func NewCmdVersion(ioStreams clioption.IOStreams) *cobra.Command {
	o := NewOptions(ioStreams)

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(o.Out, "Version: %s\n", CurrentVersion)
		},
	}

	return cmd
}
