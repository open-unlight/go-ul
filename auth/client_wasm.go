package auth

import (
	"github.com/open-unlight/go-ul/protocol"
	"github.com/open-unlight/go-ul/wasm"
	"strconv"
	"syscall/js"
)

// StartPacket ...
func (c *Client) StartPacket(this js.Value, inputs []js.Value) interface{} {
	name := inputs[0].String()
	publicKey := inputs[1].String()
	command := c.Start(name, publicKey)

	return js.TypedArrayOf(protocol.CreatePacket(command, protocol.Client))
}

// AddJSCallback ...
func (c *Client) AddJSCallback(this js.Value, inputs []js.Value) interface{} {
	argumentSize := inputs[0].Length()
	arguments := make([]string, argumentSize)

	for idx := 0; idx < argumentSize; idx++ {
		arguments[idx] = (inputs[0]).Get(strconv.Itoa(idx)).String()
	}

	idx := c.AddCallback(wasm.NewCallback(arguments, inputs[1]))

	return idx
}
