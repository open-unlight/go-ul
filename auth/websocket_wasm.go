package auth

import (
	"syscall/js"
)

// OnJSReceiveCommand ...
func (c *Client) OnJSReceiveCommand(this js.Value, inputs []js.Value) interface{} {
	buffer := []byte(inputs[0].Get("data").String())
	command := c.Decode(buffer)

	return command.Invoke()
}
