package main

import (
	"fmt"
)

type OpCode byte
type Value float64

const (
	OP_RETURN OpCode = iota
	OP_CONSTANT
	OP_NEGATE
	OP_CONSTANT_LONG
	OP_ADD
	OP_SUBTRACT
	OP_MULTIPLY
	OP_DIVIDE
)

type Chunk struct {
	code      []byte
	constants []Value
	lines     []int
}

func (c *Chunk) initChunk() {
	c.code = []byte{}
	c.constants = []Value{}
	c.lines = []int{}
}

func (c *Chunk) writeChunk(b byte, line int) {
	c.code = append(c.code, b)
	c.lines = append(c.lines, line)
}

func (c *Chunk) addConstant(v Value) int {
	c.constants = append(c.constants, v)
	return indexOf(c.constants, v)
}

func (c *Chunk) disassembleChunk(name string) {
	fmt.Printf("== %s ==\n", name)

	for offset := 0; offset < len(c.code); {
		offset = c.disassembleInstruction(offset)
	}
}

func (c *Chunk) disassembleInstruction(offset int) int {
	fmt.Printf("%04d ", offset)
	if offset > 0 && c.lines[offset] == c.lines[offset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", c.lines[offset])
	}
	instruction := c.code[offset]
	switch instruction {
	case byte(OP_RETURN):
		return c.simpleInstruction("OP_RETURN", offset)
	case byte(OP_CONSTANT):
		return c.constInstruction("OP_CONSTANT", offset)
	case byte(OP_NEGATE):
		return c.simpleInstruction("OP_NEGATE", offset)
	case byte(OP_ADD):
		return c.simpleInstruction("OP_ADD", offset)
	case byte(OP_SUBTRACT):
		return c.simpleInstruction("OP_SUBTRACT", offset)
	case byte(OP_MULTIPLY):
		return c.simpleInstruction("OP_MULTIPLY", offset)
	case byte(OP_DIVIDE):
		return c.simpleInstruction("OP_DIVIDE", offset)
	default:
		fmt.Printf("Unknow OpCode %d\n", instruction)
		return offset + 1
	}
}

func (c *Chunk) simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func (c *Chunk) constInstruction(name string, offset int) int {
	constant := c.code[offset+1]
	fmt.Printf("%s%4d", name, constant)
	printValue(c.constants[constant])
	fmt.Printf("\n")
	return offset + 2
}

func printValue(v Value) {
	fmt.Printf("%g", v)
}

/*Helper Functions*/
func indexOf[T comparable](vec []T, el T) int {
	for i, x := range vec {
		if x == el {
			return i
		}
	}
	return -1
}
