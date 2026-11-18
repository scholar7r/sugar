package buildinfo

import (
	"log/slog"
	"os"
)

var hostname string

func initHostname() {
	name, err := os.Hostname()
	if err != nil {
		slog.Error("failed to read hostname",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
	hostname = name
}
