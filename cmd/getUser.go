package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/kim-sardine/kfadmin/client/manifest"
)

// getUserCmd get all user
var getUserCmd = &cobra.Command{
	Use:   "user",
	Short: "get all user",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		cm, err := c.GetDex()
		if err != nil {
			panic(err)
		}

		dc, err := manifest.UnmarshalDexDataConfig(cm.Data["config.yaml"])
		if err != nil {
			panic(err)
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

	},
}

func init() {
	getCmd.AddCommand(getUserCmd)
}
