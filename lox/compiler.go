package main

import "fmt"

func compile(source string) {
	var scanner Scanner
	scanner.InitScanner(source)
	line := -1
	for {
		token := scanner.ScanToken()
		if token.line != line {
			fmt.Printf("%4d ", token.line)
			line = token.line
		} else {
			fmt.Printf("   | ")
		}
		fmt.Printf("  %s '%s'\n", TokToStr(token.ttpye), token.value)
		if token.ttpye == TOKEN_EOF {
			break
		}
	}
}
