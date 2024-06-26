package main

import "os"

func main() {
	var chunk Chunk
	var vm VM
	argv := os.Args
	argc := len(argv)
	if argc == 1 {
		vm.repl()
	}
	vm.initVM()
	chunk.initChunk()

	// chunk.disassembleChunk("test")
	vm.interpret(&chunk)
}
