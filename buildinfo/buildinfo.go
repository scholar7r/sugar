// SPDX-License-Identifier: GPLv2
// Copyright (c) 2026 scholar7r.

// Package buildinfo provides a minimal buildinfo printer.
//
// The buildinfo package do not provides any function to call, it has an init
// function, so that when this package has been imported, the init function
// will be executed, use slog to print the buildinfo.
package buildinfo

import (
	"log/slog"
	"os"
	"runtime"
	"runtime/debug"
)

//nolint:gochecknoinits // this init function used to print buildinfo
func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	hostname := mustHostname()

	//nolint:sloglint // use slog to print buildinfo
	slog.Info(
		"application start",
		slog.String("hostname", hostname),
		slog.String("go", info.GoVersion),
		slog.Int("cpu", runtime.NumCPU()),
		slog.String("os", runtime.GOOS),
		slog.String("arch", runtime.GOARCH),
		slog.String("vcs", info.Main.Version),
	)
}

// mustHostname exits the program if hostname retrieval fails (fail-fast).
func mustHostname() string {
	name, err := os.Hostname()
	if err != nil {
		//nolint:sloglint // use slog to print buildinfo
		slog.Error(
			"failed to get hostname",
			slog.String("error", err.Error()),
		)

		os.Exit(1)
	}

	return name
}
