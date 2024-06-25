package main

func main() {
	var chunk Chunk
	chunk.initChunk()
	chunk.writeChunk(byte(OP_RETURN), 123)
	chunk.writeChunk(byte(OP_RETURN), 124)
	index := chunk.addConstant(10)
	chunk.writeChunk(byte(OP_CONSTANT), 124)
	chunk.writeChunk(byte(index), 124)
	chunk.disassembleChunk("test")
}
