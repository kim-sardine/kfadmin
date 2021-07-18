package util

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func CkeckErr(err error) {
	if err == nil {
		return
	}

	fmt.Fprint(os.Stderr, err.Error()+"\n")
	os.Exit(1)
}

func AddRestartDexFlag(cmd *cobra.Command, restartDex *bool) {
	cmd.Flags().BoolVarP(restartDex, "restart-dex", "r", false, "Restart dex deployment to reflect changes")
}
