package main

import (
	"fmt"
	"os"
)

func main() {
	var chunk Chunk
	var vm VM
	argv := os.Args
	argc := len(argv)
	if argc == 1 {
		vm.Repl()
	} else if argc == 2 {
		vm.RunFile(argv[1])
	} else {
		fmt.Fprintf(os.Stderr, "Usage: lox [path]\n")
		os.Exit(64)
	}
	vm.initVM()
	chunk.initChunk()

	// chunk.disassembleChunk("test")
}
