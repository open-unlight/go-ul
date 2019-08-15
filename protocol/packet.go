package protocol

import (
	"bytes"
	"encoding/binary"
)

// CreatePacket ...
func CreatePacket(c Command, cType Type) []byte {
	buffer := new(bytes.Buffer)
	command := c.Encode()

	switch cType {
	case Server:
		binary.Write(buffer, binary.BigEndian, uint32(len(command)))
		break
	case Client:
		binary.Write(buffer, binary.BigEndian, uint16(len(command)))
	}

	binary.Write(buffer, binary.BigEndian, command)
	binary.Write(buffer, binary.BigEndian, []byte("\n"))

	return buffer.Bytes()
}

// DecodePacket ...
func DecodePacket(buffer []byte, arguments []string, cType Type) []interface{} {
	var decodedArguments []interface{}

	idx := 4
	if cType == Server {
		idx += 2
	}

	for _, argument := range arguments {
		switch argument {
		case "string":
			size := int(binary.BigEndian.Uint32(buffer[idx : idx+4]))
			decodedArguments = append(decodedArguments, string(buffer[idx+4:idx+4+size]))
			idx = idx + 4 + size
			break
			// TODO: Support more types
		}
	}

	return decodedArguments
}
