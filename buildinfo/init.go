package buildinfo

import (
	"fmt"
	"log/slog"
)

func init() {
	initBuildInfo()
	initHostname()
	initAppName()

	slog.Info(
		fmt.Sprintf("application %s is staring", appName),
		slog.Group(
			"svc",
			slog.String("version", info.Main.Version),
			slog.String("path", info.Main.Path),
		),
	)

	slog.Info(
		"runtime info",
		slog.String("hostname", hostname),
		slog.String("go_version", info.GoVersion),
	)
}
