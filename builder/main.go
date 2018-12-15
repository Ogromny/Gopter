package main

import (
	. "../utils"
	"io/ioutil"
	"os"
)

func main() {
	// Read the input file
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil || len(input) == 0 {
		panic(err)
	}

	// Read the stub file
	stub, err := ioutil.ReadFile("stub.elf")
	if err != nil || len(stub) == 0 {
		panic(err)
	}

	// Create the final bytes array
	encrypted := append(stub, append([]byte(Delimiter), Encrypt(input)...)...)

	// Write it
	err = ioutil.WriteFile("encrypted.elf", encrypted, 0755)
	if err != nil {
		panic(err)
	}
}
