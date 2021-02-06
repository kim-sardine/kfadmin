package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
)

// deleteProfileCmd delete kubeflow profile
var deleteProfileCmd = &cobra.Command{
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

func init() {
	deleteCmd.AddCommand(deleteProfileCmd)

	deleteProfileCmd.Flags().SortFlags = false
	deleteProfileCmd.Flags().StringP("profile", "p", "", "Profile name")
	deleteProfileCmd.MarkFlagRequired("profile")

}
