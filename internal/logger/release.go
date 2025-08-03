//go:build !debug

package logger

import (
	"log/slog"
	"os"
)

type releaseLogger struct {
	slog *slog.Logger
}

func newLogger() Logger {
	return &releaseLogger{
		slog: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo, // in prod, the logger will only show Info logs and above (Warn, Error)
		})),
	}
}

func (r *releaseLogger) Debug(msg string, args ...any) {
	// No-op (intentionally empty for prod)
}

func (r *releaseLogger) Info(msg string, args ...any) {
	r.slog.Info(msg, args...)
}

func (r *releaseLogger) Error(msg string, args ...any) {
	r.slog.Error(msg, args...)
}

func (r *releaseLogger) Enabled() bool {
	return false
}
