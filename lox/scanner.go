package main

type Scanner struct {
	source  string
	start   int
	current int
	line    int
}

func (s *Scanner) InitScanner(source string) {
	s.source = source
	s.start = 0
	s.current = 0
	s.line = 1
}

func (s *Scanner) ScanToken() Token {
	return Token{}
}
