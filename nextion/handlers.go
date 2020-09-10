package nextion

import (
	"encoding/hex"
	"log"
)

func handlerError(err error, t string) {
	switch t {
	case "panic":
		if err != nil {
			log.Panic(err)
		}
	default:
		if err != nil {
			log.Fatal(err)
		}
	}
}

func stringToHexBytes(s string) []byte {
	hx := hex.EncodeToString([]byte(s))

	decoded, err := hex.DecodeString(hx)
	handlerError(err, "")

	for _, b := range END {
		decoded = append(decoded, b)
	}

	return decoded
}
