package get

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
)

type GetDexUserOptions struct {
	KfClient *client.KfClient
	clioption.IOStreams
}

// NewGetDexUserOptions returns initialized Options
func NewGetDexUserOptions(c *client.KfClient, ioStreams clioption.IOStreams) *GetDexUserOptions {
	return &GetDexUserOptions{
		KfClient:  c,
		IOStreams: ioStreams,
	}

}

func NewCmdGetDexUsers(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewGetDexUserOptions(c, ioStreams)

	cmd := &cobra.Command{
		Use:   "dex-users",
		Short: "Print all dex users",
		Long:  `TBU`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run())
		},
	}

	return cmd
}

func (o *GetDexUserOptions) Run() error {
	cm, err := o.KfClient.GetDex()
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
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "User Email"})
	t.AppendRows(row)
	t.AppendFooter(table.Row{"Total", len(dc.StaticPasswords)})
	t.Render()

	return nil
}
