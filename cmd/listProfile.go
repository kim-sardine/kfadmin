package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listProfileCmd create kubeflow profile
var listProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "print kubeflow profile",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		profileList, err := c.GetProfileList()
		if err != nil {
			panic(err)
		}

		if len(profileList.Items) == 0 {
			fmt.Println("Kubeflow Profile does not exist")
			return
		}

		row := make([]table.Row, len(profileList.Items))
		for i, profile := range profileList.Items {
			row = append(row, table.Row{i + 1, profile.ObjectMeta.Name, profile.Spec.Owner.Name})
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Profile Name", "Owner's email"})
		t.AppendRows(row)
		t.Render()
	},
}

func init() {
	listCmd.AddCommand(listProfileCmd)
}
