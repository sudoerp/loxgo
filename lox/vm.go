package main

import (
	"fmt"
)

type InterpretResult int

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

type VM struct {
	chunk    *Chunk
	ip       int
	stack    []Value
	stackTop int
}

func (vm *VM) initVM() {
	vm.resetStack()
}

func (vm *VM) resetStack() {
	vm.stackTop = -1
}

func (vm *VM) repl() {}

func (vm *VM) push(v Value) {
	vm.stack = append(vm.stack, v)
	vm.stackTop++
}

func (vm *VM) pop() Value {
	val := vm.stack[len(vm.stack)-1]
	vm.stackTop--
	vm.stack = vm.stack[0 : vm.stackTop+1]
	return val
}

func (vm *VM) interpret(chunk *Chunk) InterpretResult {
	vm.chunk = chunk
	vm.ip = 0
	return vm.Run()
}

func (vm *VM) Run() InterpretResult {

	for {
		vm.chunk.disassembleInstruction(vm.ip)
		instruction := vm.chunk.code[vm.ip]
		vm.ip++

		switch instruction {

		case byte(OP_CONSTANT):
			constant := vm.chunk.code[vm.ip]
			vm.push(vm.chunk.constants[constant])
			vm.ip++
			fmt.Printf("CONSTANT VALUE: ")
			printValue(vm.chunk.constants[constant])
			fmt.Printf("\n")

		case byte(OP_RETURN):
			printValue(vm.pop())
			println()
			return INTERPRET_OK

		case byte(OP_NEGATE):
			vm.push(-vm.pop())

		case byte(OP_ADD):
			b := vm.pop()
			a := vm.pop()
			vm.push(a + b)

		case byte(OP_SUBTRACT):
			b := vm.pop()
			a := vm.pop()
			vm.push(a - b)

		case byte(OP_MULTIPLY):
			b := vm.pop()
			a := vm.pop()
			vm.push(a * b)

		case byte(OP_DIVIDE):
			b := vm.pop()
			a := vm.pop()
			vm.push(a / b)
		default:
			fmt.Printf("Invalid")
			return INTERPRET_RUNTIME_ERROR
		}

		fmt.Printf("       ")

		for i := 0; i < vm.stackTop; i++ {
			fmt.Printf("[")
			printValue(vm.stack[i])
			fmt.Printf("]")
		}

		fmt.Println()
	}
}
