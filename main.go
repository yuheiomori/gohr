package main

import (
	"os"
	"syscall"

	"code.google.com/p/go.crypto/ssh/terminal"
)

func main() {

	width, _, _ := terminal.GetSize(syscall.Stdin)

	args := os.Args[1:]
	if len(args) == 0 {
		args = append(args, "#")
	}

	for _, arg := range args {
		line := make([]byte, 0)
		bytes := []byte(arg)
		for len(line) < width {
			line = append(line, bytes...)
		}
		os.Stdout.Write(line[:width])
	}
}
