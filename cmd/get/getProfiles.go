package get

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/spf13/cobra"
)

type GetProfileOptions struct {
	KfClient *client.KfClient
	clioption.IOStreams
}

// NewGetProfileOptions returns initialized Options
func NewGetProfileOptions(c *client.KfClient, ioStreams clioption.IOStreams) *GetProfileOptions {
	return &GetProfileOptions{
		KfClient:  c,
		IOStreams: ioStreams,
	}

}

func NewCmdGetProfiles(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewGetProfileOptions(c, ioStreams)

	cmd := &cobra.Command{
		Use:   "profiles",
		Short: "Print all kubeflow profiles",
		Long:  `TBU`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run())
		},
	}

	return cmd
}

func (o *GetProfileOptions) Run() error {
	profiles, err := o.KfClient.GetProfiles()
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
