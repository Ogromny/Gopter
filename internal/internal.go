package internal

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
)

var delimiter = []byte("ThisIsTheDelimiter")

func init() {
	delimiter = Encode(delimiter)
}

func Decode(data []byte) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	if _, err := base64.StdEncoding.Decode(dst, data); err != nil {
		return nil, err
	}
	return dst, nil
}

func Encode(data []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	return dst
}

func Bind(stub, input []byte) []byte {
	return append(append(stub, []byte(delimiter)...), input...)
}

func Extract(data []byte) []byte {
	if pos := bytes.LastIndex(data, delimiter); pos != -1 {
		return data[pos+len(delimiter):]
	}
	return nil
}

func Perror(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}
