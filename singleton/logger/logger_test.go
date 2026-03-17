package logger_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/scholar7r/sugar/singleton/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestGetLogger(t *testing.T) {
	log := logger.Get()
	if log == nil {
		t.Fatal("expected logger to be initialized, got nil")
	}

	var b bytes.Buffer
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(&b),
		zapcore.DebugLevel,
	)
	testLogger := zap.New(core)
	defer testLogger.Sync()

	testLogger.Info("test message")
	logged := b.String()

	if !strings.Contains(logged, "test message") {
		t.Errorf("expected log to contain 'test message', got: %s", logged)
	}
}

func TestSingletonLogger(t *testing.T) {
	l1 := logger.Get()
	l2 := logger.Get()

	if l1 != l2 {
		t.Errorf("expected same logger instance, got different instances")
	}
}
