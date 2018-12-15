package main

import (
	. "../utils"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Read himself
	bytes, err := ioutil.ReadFile(os.Args[0])
	if err != nil || len(bytes) == 0 {
		panic(err)
	}

	// Find the delimiter position
	pos := strings.LastIndex(string(bytes), Delimiter)

	// If it does not exist it means that
	// the stub is not "finalized"
	if pos == -1 {
		os.Exit(1)
	}

	// Append the delimiter length
	pos += len(Delimiter)

	// Get the File stored in stub
	file := bytes[pos:]

	// Write it
	err = ioutil.WriteFile("decrypted.elf", Decrypt(file), 0644)
	if err != nil {
		panic(err)
	}
}
