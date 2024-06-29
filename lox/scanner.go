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
	s.SkipWhitespace()
	s.start = s.current
	if s.IsAtEnd() {
		return s.MakeToken(TOKEN_EOF)
	}

	c := s.Advance()
	if IsDigit(c) {
		return s.Number()
	}
	if IsAlpha(c) {
		return s.Indentifier()
	}
	switch c {
	case '(':
		return s.MakeToken(TOKEN_LEFT_PAREN)
	case ')':
		return s.MakeToken(TOKEN_RIGHT_PAREN)
	case '{':
		return s.MakeToken(TOKEN_LEFT_BRACE)
	case '}':
		return s.MakeToken(TOKEN_RIGHT_BRACE)
	case ';':
		return s.MakeToken(TOKEN_SEMICOLON)
	case ',':
		return s.MakeToken(TOKEN_COMMA)
	case '.':
		return s.MakeToken(TOKEN_DOT)
	case '-':
		return s.MakeToken(TOKEN_MINUS)
	case '+':
		return s.MakeToken(TOKEN_PLUS)
	case '/':
		return s.MakeToken(TOKEN_SLASH)
	case '*':
		return s.MakeToken(TOKEN_STAR)
	case '!':
		if s.Match('=') {
			return s.MakeToken(TOKEN_BANG_EQUAL)
		} else {
			return s.MakeToken(TOKEN_BANG)
		}
	case '=':
		if s.Match('=') {
			return s.MakeToken(TOKEN_EQUAL_EQUAL)
		} else {
			return s.MakeToken(TOKEN_EQUAL)
		}
	case '<':
		if s.Match('=') {
			return s.MakeToken(TOKEN_LESS_EQUAL)
		} else {
			return s.MakeToken(TOKEN_LESS)
		}
	case '>':
		if s.Match('=') {
			return s.MakeToken(TOKEN_GREATER_EQUAL)
		} else {
			return s.MakeToken(TOKEN_GREATER)
		}
	case '#':
		for s.Peek() != '\n' || s.IsAtEnd() {
			s.Advance()
		}
	case '"':
		return s.String()
	}

	return Token{}
}

func (s *Scanner) Indentifier() Token {
	for IsAlpha(s.Peek()) || IsDigit(s.Peek()) {
		s.Advance()
	}
	return s.MakeToken(s.IdentifierType())
}

func (s *Scanner) IdentifierType() TokenType {
	ident := s.source[s.start:s.current]
	switch ident {
	case "false":
		return TOKEN_FALSE
	case "and":
		return TOKEN_AND
	case "or":
		return TOKEN_OR
	case "class":
		return TOKEN_CLASS
	case "if":
		return TOKEN_IF
	case "else":
		return TOKEN_ELSE
	case "fn":
		return TOKEN_FUN
	case "nil":
		return TOKEN_NIL
	case "print":
		return TOKEN_PRINT
	case "return":
		return TOKEN_RETURN
	case "super":
		return TOKEN_SUPER
	case "this":
		return TOKEN_THIS
	case "true":
		return TOKEN_TRUE
	case "let":
		return TOKEN_LET
	case "while":
		return TOKEN_WHILE
	}
	return TOKEN_IDENTIFIER
}

func (s *Scanner) Number() Token {
	for IsDigit(s.Peek()) {
		s.Advance()
	}
	if s.Peek() == '.' && IsDigit(s.PeekNext()) {
		s.Advance()

		for IsDigit(s.Peek()) {
			s.Advance()
		}
	}
	return s.MakeToken(TOKEN_NUMBER)
}

func (s *Scanner) String() Token {
	for s.Peek() != '"' && !s.IsAtEnd() {
		if s.Peek() == '\n' {
			s.line++
		}
		s.Advance()
	}
	if s.IsAtEnd() {
		return s.ErrorToken("Unterminated String!")
	}
	s.Advance()
	return s.MakeToken(TOKEN_STRING)
}

func (s *Scanner) Match(expected rune) bool {
	if s.IsAtEnd() {
		return false
	}
	if rune(s.source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) SkipWhitespace() {
	for {
		c := s.Peek()
		switch c {
		case ' ':
			s.Advance()
		case '\r':
			s.Advance()
		case '\t':
			s.Advance()
		case '\n':
			s.line++
			s.Advance()
		default:
			return
		}
	}
}

func (s *Scanner) PeekNext() rune {
	if s.IsAtEnd() || s.current+1 >= len(s.source) {
		return 0
	}
	return rune(s.source[s.current+1])
}

func (s *Scanner) Peek() rune {
	if s.IsAtEnd() {
		return 0
	}
	return rune(s.source[s.current])
}

func (s *Scanner) Advance() rune {
	s.current++
	return rune(s.source[s.current-1])
}

func (s *Scanner) MakeToken(ttype TokenType) Token {
	return Token{
		ttpye: ttype,
		value: s.source[s.start:s.current],
		line:  s.line,
	}
}

func (s *Scanner) ErrorToken(msg string) Token {
	return Token{
		ttpye: TOKEN_ERROR,
		value: msg,
		line:  s.line,
	}
}

func (s *Scanner) IsAtEnd() bool {
	return s.current >= len(s.source)
}

func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c == '_')
}
