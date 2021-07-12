package delete

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/spf13/cobra"
)

func NewCmdDeleteStaticUser(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:     "staticuser",
		Aliases: []string{"static-user"},
		Short:   "Delete dex static user",
		Long:    `Delete dex static user`,
		Run: func(cmd *cobra.Command, args []string) {
			email, _ := cmd.Flags().GetString("email")

			// Check if user exists
			cm, err := c.GetDex()
			if err != nil {
				panic(err)
			}

			originalData := cm.Data["config.yaml"]
			dc, err := manifest.UnmarshalDexDataConfig(originalData)
			if err != nil {
				panic(err)
			}

			users := dc.StaticPasswords
			userIdx := -1
			for i, user := range users {
				if user.Email == email {
					userIdx = i
					break
				}
			}
			if userIdx < 0 {
				panic(fmt.Errorf("user with email '%s' does not exist", email))
			}

			// Delete if exists
			dc.StaticPasswords = append(dc.StaticPasswords[:userIdx], dc.StaticPasswords[userIdx+1:]...)
			cm.Data["config.yaml"], err = manifest.MarshalDexDataConfig(dc)
			if err != nil {
				panic(err)
			}

			err = c.UpdateDex(cm)
			if err != nil {
				panic(err)
			}

			err = c.RestartDexDeployment(originalData)
			if err != nil {
				panic(err)
			}
			fmt.Printf("static user '%s' deleted\n", email)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringP("email", "e", "", "User email")
	cmd.MarkFlagRequired("email")

	return cmd
}
