package cmd

import (
	"github.com/spf13/cobra"
)

// updateUserCmd update user info
var updateUserCmd = &cobra.Command{
	Use:   "user",
	Short: "update user information",
	Long:  `TBU`,
}

func init() {
	updateCmd.AddCommand(updateUserCmd)
}
