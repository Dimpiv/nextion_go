package nextion

import (
	"encoding/hex"
	"log"
)

func stringToHexBytes(s string) []byte {
	hx := hex.EncodeToString([]byte(s))

	decoded, err := hex.DecodeString(hx)
	if err != nil {
		log.Println(err)
	}

	for _, b := range EndNextionMessage {
		decoded = append(decoded, b)
	}

	return decoded
}
