package delete

import (
	"fmt"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/kim-sardine/kfadmin/clioption"
	"github.com/kim-sardine/kfadmin/cmd/util"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/errors"
)

type DeleteProfileOptions struct {
	Profile string

	clioption.IOStreams
}

func NewDeleteProfileOptions(ioStreams clioption.IOStreams) *DeleteProfileOptions {
	return &DeleteProfileOptions{
		IOStreams: ioStreams,
	}
}

func NewCmdDeleteProfile(c *client.KfClient, ioStreams clioption.IOStreams) *cobra.Command {
	o := NewDeleteProfileOptions(ioStreams)

	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Delete kubeflow profile",
		Long:  `Delete kubeflow profile`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CkeckErr(o.Run(c, cmd))
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&o.Profile, "profile", "p", "", "Profile name")
	cmd.MarkFlagRequired("profile")

	return cmd
}

func (o *DeleteProfileOptions) Run(c *client.KfClient, cmd *cobra.Command) error {

	profileName, _ := cmd.Flags().GetString("profile")

	if _, err := c.GetProfile(profileName); err != nil {
		if errors.IsNotFound(err) {
			return fmt.Errorf("kubeflow profile '%s' does not exist", profileName)

		} else {
			return err
		}
	}

	if err := c.DeleteProfile(profileName); err != nil {
		return err
	}

	fmt.Fprintf(o.Out, "Profile '%s' deleted\n", profileName)

	return nil
}
