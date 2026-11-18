package buildinfo

import (
	"log/slog"
	"os"
	"runtime/debug"
)

var info *debug.BuildInfo

func initBuildInfo() {
	v, ok := debug.ReadBuildInfo()
	if !ok {
		slog.Error("failed to read build info")
		os.Exit(1)
	}
	info = v
}
