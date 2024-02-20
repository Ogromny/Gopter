package main

import (
	"gopter/internal"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		internal.Perror("usage: %s input-file\n", os.Args[0])
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		internal.Perror("cannot read %s: %s\n", os.Args[1], err)
	}

	stub, err := os.ReadFile("stub")
	if err != nil {
		internal.Perror("cannot read stub: %s\n", err)
	}

	if err := os.WriteFile("binded", internal.Bind(stub, input), 0755); err != nil {
		internal.Perror("cannot write binded: %s\n", err)
	}
}
