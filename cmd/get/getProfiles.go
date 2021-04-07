package get

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/spf13/cobra"
)

type GetProfileOptions struct {
	clioption.IOStreams
}

// NewGetProfileOptions returns initialized Options
func NewGetProfileOptions(ioStreams clioption.IOStreams) *GetProfileOptions {
	return &GetProfileOptions{
		IOStreams: ioStreams,
	}

}

func NewCmdGetProfiles(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewGetProfileOptions(ioStreams)

	cmd := &cobra.Command{
		Use:   "profiles",
		Short: "print kubeflow profiles",
		Long:  `TBU`,
		Run: func(cmd *cobra.Command, args []string) {
			profiles, err := c.GetProfiles()
			if err != nil {
				panic(err)
			}

			if len(profiles.Items) == 0 {
				fmt.Fprintf(o.Out, "Kubeflow Profile does not exist")
				return
			}

			row := make([]table.Row, len(profiles.Items))
			for i, profile := range profiles.Items {
				row = append(row, table.Row{i + 1, profile.ObjectMeta.Name, profile.Spec.Owner.Name})
			}

			t := table.NewWriter()
			t.SetOutputMirror(o.Out)
			t.AppendHeader(table.Row{"#", "Profile Name", "Owner's email"})
			t.AppendRows(row)
			t.Render()
		},
	}

	return cmd
}
