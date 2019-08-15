package protocol

import (
	"bytes"
	"encoding/binary"
)

// Command ...
type Command struct {
	ID        uint16
	Type      Type
	Arguments []interface{}
	Callback  *Callback
}

// Encode ...
func (c *Command) Encode() []byte {
	// TODO: Encode Server
	return encodeAsClient(c.ID, c.Arguments)
}

// Invoke ...
func (c *Command) Invoke() interface{} {
	if c.Callback != nil {
		return (*c.Callback).Invoke(c.Arguments...)
	}

	return nil
}

func encodeAsClient(id uint16, arguments []interface{}) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, id)

	for _, argument := range arguments {
		switch v := argument.(type) {
		case string:
			binary.Write(buffer, binary.BigEndian, uint16(len(v)))
			binary.Write(buffer, binary.BigEndian, []byte(v))
			break
			// TODO: Support more types
		}
	}

	return buffer.Bytes()
}
