package admin

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// ListProfile print kubeflow profile
func ListProfile() {

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
}
