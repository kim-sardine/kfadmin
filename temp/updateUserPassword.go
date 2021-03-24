package cmd

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/spf13/cobra"
)

// updateUserPasswordCmd udpate staticPassword's password
var updateUserPasswordCmd = &cobra.Command{
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
		hashedPassword, err := hashPassword(password)
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

func init() {
	updateUserCmd.AddCommand(updateUserPasswordCmd)

	updateUserPasswordCmd.Flags().SortFlags = false
	updateUserPasswordCmd.Flags().StringP("email", "e", "", "User email")
	updateUserPasswordCmd.MarkFlagRequired("email")
	updateUserPasswordCmd.Flags().StringP("password", "p", "", "User password")
	updateUserPasswordCmd.MarkFlagRequired("password")
}
