package get

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
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
		Use:     "profiles",
		Aliases: []string{"profile"},
		Short:   "Print all kubeflow profiles",
		Long:    `Print all kubeflow profiles`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run(c, cmd))
		},
	}

	return cmd
}

func (o *GetProfileOptions) Run(c *client.KfClient, cmd *cobra.Command) error {
	profiles, err := c.GetProfiles()
	if err != nil {
		return err
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

	return nil
}
