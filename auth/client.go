package auth

import (
	"encoding/binary"
	"github.com/open-unlight/go-ul/protocol"
)

// Client is AuthServer Client
type Client struct {
	callbacks []protocol.Callback
	// TODO: Register Handler
}

// NewClient will create a AuthServer Client
func NewClient() *Client {
	return &Client{}
}

// Start is Start Auth
func (c *Client) Start(name string, cPubKey string) protocol.Command {
	return protocol.Command{ID: 1, Type: protocol.Client, Arguments: []interface{}{name, cPubKey}}
}

// Decode is decode server command
func (c *Client) Decode(buffer []byte) protocol.Command {
	commandID := binary.BigEndian.Uint16(buffer[4:6])

	// TODO: Command must correct registered
	callback := c.callbacks[commandID]
	arguments := protocol.DecodePacket(buffer, callback.Arguments(), protocol.Server)

	return protocol.Command{ID: commandID, Type: protocol.Server, Arguments: arguments, Callback: &callback}
}

// AddCallback ...
func (c *Client) AddCallback(cb protocol.Callback) int {
	idx := len(c.callbacks)
	c.callbacks = append(c.callbacks, cb)

	return idx
}
