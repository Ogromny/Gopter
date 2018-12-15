package utils

import (
	"encoding/base64"
	"fmt"
)

var Delimiter = "ThisIsTheDelimiter"

func Decrypt(data []byte) []byte {
	// Create array of decoded length
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))

	// Decode the data in dst
	n, err := base64.StdEncoding.Decode(dst, data)
	if err != nil || n + 1 != base64.StdEncoding.DecodedLen(len(data)) {
		panic(err)
	}

	fmt.Println("n: ", n)
	fmt.Println("decodelen: ", base64.StdEncoding.DecodedLen(len(data)))

	return dst
}

func Encrypt(data []byte) []byte {
	// Create array of encoded length
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))

	// Encode data in dst
	base64.StdEncoding.Encode(dst, data)

	return dst
}
