package wasm

import (
	"syscall/js"
)

// Callback ...
type Callback struct {
	Func      js.Value
	arguments []string
}

// NewCallback ...
func NewCallback(arguments []string, cb js.Value) *Callback {
	return &Callback{cb, arguments}
}

// Arguments ...
func (c *Callback) Arguments() []string {
	return c.arguments
}

// Invoke ..
func (c *Callback) Invoke(args ...interface{}) interface{} {
	return c.Func.Invoke(args...)
}
