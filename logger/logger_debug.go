//+build debug,!wasm

package logger

import (
	"log"
	"os"
)

var (
	infoLog  *log.Logger
	debugLog *log.Logger
	errorLog *log.Logger
)

func init() {
	infoLog = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile)
	debugLog = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
	errorLog = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)
}

// Info ...
func Info(format string, args ...interface{}) {
	infoLog.Printf(format, args...)
}

// Debug ...
func Debug(format string, args ...interface{}) {
	debugLog.Printf(format, args...)
}

// Error ...
func Error(format string, args ...interface{}) {
	errorLog.Printf(format, args...)
}
