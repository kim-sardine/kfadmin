package get

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/kim-sardine/kfadmin/manifest"
)

type GetStaticUserOptions struct {
	clioption.IOStreams
}

// NewGetStaticUserOptions returns initialized Options
func NewGetStaticUserOptions(ioStreams clioption.IOStreams) *GetStaticUserOptions {
	return &GetStaticUserOptions{
		IOStreams: ioStreams,
	}

}

func NewCmdGetStaticUsers(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewGetStaticUserOptions(ioStreams)

	cmd := &cobra.Command{
		Use:     "staticusers",
		Aliases: []string{"staticuser", "static-users"},
		Short:   "Print all dex static users",
		Long:    `Print all dex static users`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run(c, cmd))
		},
	}

	return cmd
}

func (o *GetStaticUserOptions) Run(c *client.KfClient, cmd *cobra.Command) error {
	cm, err := c.GetDexConfigMap()
	if err != nil {
		return err
	}

	dc, err := manifest.UnmarshalDexDataConfig(cm.Data["config.yaml"])
	if err != nil {
		return err
	}

	row := make([]table.Row, len(dc.StaticPasswords))
	for i, user := range dc.StaticPasswords {
		row = append(row, table.Row{i + 1, user.Email})
	}

	t := table.NewWriter()
	t.SetOutputMirror(o.Out)
	t.AppendHeader(table.Row{"#", "User Email"})
	t.AppendRows(row)
	t.AppendFooter(table.Row{"Total", len(dc.StaticPasswords)})
	t.Render()

	return nil
}
