package buildinfo

import (
	"os"
	"path/filepath"
)

var appName string

func initAppName() {
	name := "unknown"
	if len(os.Args) > 0 {
		name = filepath.Base(os.Args[0])
	}
	appName = name
}
