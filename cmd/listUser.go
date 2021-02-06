package cmd

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/spf13/cobra"
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
		for i, user := range dc.StaticPasswords {
			fmt.Println(i+1, user.Email)
		}
		fmt.Printf("Total : %d users\n", len(dc.StaticPasswords))
	},
}

func init() {
	listCmd.AddCommand(listUserCmd)
}
