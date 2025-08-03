//go:build debug

package logger

import (
	"log/slog"
	"os"
)

type debugLogger struct {
	slog *slog.Logger
}

/*
Initializes logger with detailed output, config used:
  - Level: slog.LevelDebug (-4) means the logger will show logs that are above -4 (since this is the lowest level, basically all of them)
  - AddSource shows the line in the source code where the log is
*/
func newLogger() Logger {
	return &debugLogger{
		slog: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	}
}

func (d *debugLogger) Debug(msg string, args ...any) {
	d.slog.Debug(msg, args...)
}

func (d *debugLogger) Info(msg string, args ...any) {
	d.slog.Info(msg, args...)
}

func (d *debugLogger) Error(msg string, args ...any) {
	d.slog.Error(msg, args...)
}

func (d *debugLogger) Enabled() bool {
	return true
}
