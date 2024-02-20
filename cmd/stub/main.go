package main

import (
	"gopter/internal"
	"os"
	"os/exec"
)

func main() {
	data, err := os.ReadFile(os.Args[0])
	if err != nil {
		internal.Perror("cannot read %s: %s\n", os.Args[0], err)
	}

	file := internal.Extract(data)
	if file == nil {
		return
	}

	tmpFile, err := os.CreateTemp("", "*")
	if err != nil {
		internal.Perror("cannot create temp file: %s\n", err)
	}
	defer func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}()

	if _, err := tmpFile.Write(file); err != nil {
		internal.Perror("cannot write temp file: %s\n", err)
	}
	if err := tmpFile.Chmod(0755); err != nil {
		internal.Perror("cannot change temp file mod: %s", err)
	}
	tmpFile.Close()

	cmd := exec.Command(tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		internal.Perror("cannot execute temp file: %s\n", err)
	}
}
