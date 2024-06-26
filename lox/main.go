package main

func main() {
	var chunk Chunk
	var vm VM
	vm.initVM()
	chunk.initChunk()
	index := chunk.addConstant(10)
	chunk.writeChunk(byte(OP_CONSTANT), 124)
	chunk.writeChunk(byte(index), 124)
	chunk.writeChunk(byte(OP_NEGATE), 124)
	chunk.writeChunk(byte(OP_RETURN), 124)
	// chunk.disassembleChunk("test")
	vm.interpret(&chunk)
}
