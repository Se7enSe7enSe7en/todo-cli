package logger

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Enabled() bool
}

// Global logger instance
var Default Logger

func Initialize() {
	Default = newLogger()
}

func Debug(msg string, args ...any) {
	Default.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	Default.Info(msg, args...)
}

func Error(msg string, args ...any) {
	Default.Error(msg, args)
}

func Enabled() bool {
	return Default.Enabled()
}
