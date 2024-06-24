package lox

type OpCode byte

const (
	OP_RETURN OpCode = iota
)

type Chunk struct {
	code []byte
}

func (c *Chunk) initChunk() {
	c.code = []byte{}
}

func (c *Chunk) writeChunk(b byte) {
	c.code = append(c.code, b)
}
