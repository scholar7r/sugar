// SPDX-License-Identifier: GPLv2
// Copyright (c) 2026 scholar7r.

// Package logger provides a global, ready-to-use zap.Logger instance.
//
// It uses a singleton pattern to ensure that the logger is initialized only once.
// The logger outputs to the console with colored level encoding, ISO8601 timestamps,
// and includes caller information and stacktraces for errors.
package logger

import (
	"os"

	"github.com/scholar7r/sugar/singleton"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var global = singleton.New(load)

// Get returns the global *zap.Logger instance.
//
// The logger is lazily initialized on first use and is safe for concurrent use.
func Get() *zap.Logger {
	return *global.Get()
}

// load initializes a zap.Logger with a console encoder, colored levels,
// ISO8601 timestamps, and stacktraces for error-level logs.
func load() *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	return zap.New(
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}
