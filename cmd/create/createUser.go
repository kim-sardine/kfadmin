package create

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/client/manifest"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/spf13/cobra"
)

// NewCmdCreateUser TBU
func NewCmdCreateUser(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "user",
		Short: "create new user",
		Long:  `TBU`,
		Run: func(cmd *cobra.Command, args []string) {
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")
			// TODO: into shared util
			// Checkout https://github.com/kubernetes/kubectl/tree/master/pkg/cmd/create
			restartDex, _ := cmd.Flags().GetBool("restart-dex")

			username, err := util.GetUsernameFromEmail(email)
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

			hashedPassword, err := util.HashPassword(password)
			if err != nil {
				panic(err)
			}

			newUser := manifest.StaticPassword{
				Email:    email,
				Hash:     hashedPassword,
				Username: username,
				UserID:   util.GetUniqueUUID(uuids),
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

			// TODO: into shared util
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

	cmd.Flags().SortFlags = false
	cmd.Flags().StringP("email", "e", "", "User email")
	cmd.MarkFlagRequired("email")
	cmd.Flags().StringP("password", "p", "", "User password")
	cmd.MarkFlagRequired("password")
	// TODO: into shared util
	cmd.Flags().BoolP("restart-dex", "r", false, "Restart dex deployment to reflect changes")

	return cmd
}
