package cmd

import (
	"os"
)

func checkStreamIsPipe() bool {
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		return true
	}

	return false
}
