//go:build debug

package logger

import (
	"fmt"
	"log"
)

func Debug(format string, args ...any) {
	fmt.Printf("[DEBUG] "+format, args...)
	fmt.Println()
}

func Error(msg string, args ...any) {
	log.Fatal("[ERROR] ", args)
}
