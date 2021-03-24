package get

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
)

func NewCmdGetProfiles(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "profile",
		Short: "print kubeflow profile",
		Long:  `TBU`,
		Run: func(cmd *cobra.Command, args []string) {
			profiles, err := c.GetProfiles()
			if err != nil {
				panic(err)
			}

			if len(profiles.Items) == 0 {
				fmt.Println("Kubeflow Profile does not exist")
				return
			}

			row := make([]table.Row, len(profiles.Items))
			for i, profile := range profiles.Items {
				row = append(row, table.Row{i + 1, profile.ObjectMeta.Name, profile.Spec.Owner.Name})
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"#", "Profile Name", "Owner's email"})
			t.AppendRows(row)
			t.Render()
		},
	}

	return cmd
}
