package cmd

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/spf13/cobra"
)

// createUserCmd create new user
var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "create new user",
	Long:  `TBU`,
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")
		// TODO: into shared utils
		// Checkout https://github.com/kubernetes/kubectl/tree/master/pkg/cmd/create
		restartDex, _ := cmd.Flags().GetBool("restart-dex")

		username, err := getUsernameFromEmail(email)
		if err != nil {
			panic(err)
		}

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

		uuids := make([]string, len(users)+1)
		for _, user := range users {
			uuids = append(uuids, user.UserID)
		}

		hashedPassword, err := hashPassword(password)
		if err != nil {
			panic(err)
		}

		newUser := manifest.StaticPassword{
			Email:    email,
			Hash:     hashedPassword,
			Username: username,
			UserID:   getUniqueUUID(uuids),
		}

		dc.StaticPasswords = append(dc.StaticPasswords, newUser)

		cm.Data["config.yaml"], err = manifest.MarshalDexDataConfig(dc)
		if err != nil {
			panic(err)
		}

		err = c.UpdateDex(cm)
		if err != nil {
			panic(err)
		}

		// TODO: into shared utils
		if restartDex {
			err = c.RestartDexDeployment(originalData)
			if err != nil {
				panic(err)
			}
			fmt.Printf("user '%s' created\n", email)
		} else {
			fmt.Printf("user '%s' added to dex ConfigMap\nTo reflext changes, please run below command\n\nkubectl rollout restart deployment dex -n auth\n\n", email)
		}
	},
}

func init() {
	createCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().SortFlags = false
	createUserCmd.Flags().StringP("email", "e", "", "User email")
	createUserCmd.MarkFlagRequired("email")
	createUserCmd.Flags().StringP("password", "p", "", "User password")
	createUserCmd.MarkFlagRequired("password")
	// TODO: into shared utils
	createUserCmd.Flags().BoolP("restart-dex", "r", false, "Restart dex deployment to reflect changes")
}
