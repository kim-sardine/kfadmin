package main

import (
	"os"

	"github.com/kim-sardine/kfadmin/cmd"
)

func main() {
	cmd := cmd.NewKfAdminCommand(os.Stdin, os.Stdout, os.Stderr)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
