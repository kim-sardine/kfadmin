package delete

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
)

func NewCmdDeleteProfile(c *client.KfClient) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "profile",
		Short: "delete kubeflow profile",
		Long:  `TBU`,
		Run: func(cmd *cobra.Command, args []string) {
			profileName, _ := cmd.Flags().GetString("profile")

			_, err := c.GetProfile(profileName)
			if err != nil {
				if errors.IsNotFound(err) {
					panic(fmt.Errorf("Kubeflow profile '%s' does not exist", profileName))

				} else {
					panic(err)
				}
			}

			err = c.DeleteProfile(profileName)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Profile '%s' deleted\n", profileName)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringP("profile", "p", "", "Profile name")
	cmd.MarkFlagRequired("profile")

	return cmd
}
