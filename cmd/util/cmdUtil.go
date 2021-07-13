package util

import (
	"fmt"
	"os"
)

func CkeckErr(err error) {
	if err == nil {
		return
	}

	fmt.Fprint(os.Stderr, err.Error()+"\n")
	os.Exit(1)
}
