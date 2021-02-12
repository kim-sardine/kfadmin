package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/kim-sardine/kfadmin/client/manifest"
)

// listUserCmd list all user
var listUserCmd = &cobra.Command{
	Use:   "user",
	Short: "list all user",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		cm, err := c.GetDex()
		if err != nil {
			panic(err)
		}

		dc := manifest.UnmarshalDexConfig(cm.Data["config.yaml"])

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

	},
}

func init() {
	listCmd.AddCommand(listUserCmd)
}
