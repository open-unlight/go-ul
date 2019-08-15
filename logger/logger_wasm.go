//+build debug,wasm

package logger

import (
	"fmt"
	"syscall/js"
)

var console js.Value

func init() {
	console = js.Global().Get("console")
}

// Info ...
func Info(format string, args ...interface{}) {
	console.Call("info", fmt.Sprintf(format, args...))
}

// Debug ...
func Debug(format string, args ...interface{}) {
	console.Call("debug", fmt.Sprintf(format, args...))
}

// Error ...
func Error(format string, args ...interface{}) {
	console.Call("error", fmt.Sprintf(format, args...))
}
