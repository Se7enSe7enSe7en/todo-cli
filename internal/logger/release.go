//go:build !debug

package logger

import "log"

func Debug(msg string, args ...any) {
	// intentionally empty function
}

func Error(msg string, args ...any) {
	log.Fatal("[ERROR] ", args)
}
