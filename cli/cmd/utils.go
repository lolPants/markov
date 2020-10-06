package cmd

import (
	"os"
)

func checkIsPipe(f *os.File) bool {
	fi, _ := f.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		return true
	}

	return false
}
