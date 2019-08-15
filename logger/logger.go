//+build !debug

package logger

// Info ...
func Info(format string, args ...interface{}) {}

// Debug ...
func Debug(format string, args ...interface{}) {}

// Error
func Error(format string, args ...interface{}) {}
