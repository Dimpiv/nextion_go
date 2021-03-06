package nextion

import (
	"encoding/hex"
)

var EndNextionMessage = []byte{0xFF, 0xFF, 0xFF} // All instructions over serial: are terminated with three bytes of 0xFF 0xFF 0xFF

const (
	ERROR   = 0x00 // Returned when instruction sent by user has failed
	SUCCESS = 0x01 // Returned when instruction sent by user was successful
)

func CheckReturnedCode(code []byte) string {
	switch code[0] {
	case ERROR:
		return "Error command (0x00)"
	case SUCCESS:
		return "Command apply (0x01)"
	default:
		return hex.EncodeToString(code)
	}
}
