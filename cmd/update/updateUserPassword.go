package update

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/spf13/cobra"
)

func NewCmdUpdateUserPassword(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "password",
		Short: "udpate dex staticPassword's password",
		Long:  `TBU`,
		Run: func(cmd *cobra.Command, args []string) {
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")

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

			// Change Password
			hashedPassword, err := util.HashPassword(password)
			if err != nil {
				panic(err)
			}

			dc.StaticPasswords[userIdx].Hash = hashedPassword
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
			fmt.Printf("user '%s' updated\n", email)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringP("email", "e", "", "User email")
	cmd.MarkFlagRequired("email")
	cmd.Flags().StringP("password", "p", "", "User password")
	cmd.MarkFlagRequired("password")

	return cmd
}
