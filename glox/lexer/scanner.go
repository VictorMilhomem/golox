package lexer

import (
	"strconv"

	"github.com/VictorMilhomem/glox/glox/utils"
)

type Scanner struct {
	source   string
	tokens   []Token
	Keywords map[string]TokenType
	start    int
	current  int
	line     int
}

func NewScanner(s string) *Scanner {
	kw := make(map[string]TokenType)
	kw["and"] = AND
	kw["class"] = CLASS
	kw["else"] = ELSE
	kw["false"] = FALSE
	kw["for"] = FOR
	kw["fun"] = FUN
	kw["if"] = IF
	kw["nil"] = NIL
	kw["or"] = OR
	kw["print"] = PRINT
	kw["return"] = RETURN
	kw["super"] = SUPER
	kw["this"] = THIS
	kw["true"] = TRUE
	kw["var"] = VAR
	kw["while"] = WHILE
	return &Scanner{
		source:   s,
		Keywords: kw,
		start:    0,
		current:  0,
		line:     1,
	}
}

func (s *Scanner) advance() string {
	newString := string(s.source[s.current])
	s.current++
	return newString
}

func (s *Scanner) _addToken(t TokenType) {
	s.addToken(t, nil)
}

func (s *Scanner) addToken(t TokenType, literal Object) {
	text := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, Token{t, text, literal, s.line})
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case "(":
		s._addToken(LEFT_PAREN)
	case ")":
		s._addToken(RIGHT_PAREN)
	case "{":
		s._addToken(LEFT_BRACE)
	case "}":
		s._addToken(RIGHT_BRACE)
	case ",":
		s._addToken(COMMA)
	case ".":
		s._addToken(DOT)
	case "-":
		s._addToken(MINUS)
	case "+":
		s._addToken(PLUS)
	case ";":
		s._addToken(SEMICOLON)
	case "*":
		s._addToken(STAR)
	case "!":
		if s.match("=") {
			s._addToken(BANG_EQUAL)
		} else {
			s._addToken(BANG)
		}
	case "=":
		if s.match("=") {
			s._addToken(EQUAL_EQUAL)
		} else {
			s._addToken(EQUAL)
		}
	case "<":
		if s.match("=") {
			s._addToken(LESS_EQUAL)
		} else {
			s._addToken(LESS)
		}
	case ">":
		if s.match("=") {
			s._addToken(GREATER_EQUAL)
		} else {
			s._addToken(GREATER)
		}
	case "/":
		if s.match("/") {
			// Single line comments
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else if s.match("*") {
			// Multi line comments
			toMatch := 1
			for !s.isAtEnd() {
				if s.peek() == "*" && s.peekNext() == "/" {
					toMatch--
				}
				if s.peek() == "/" && s.peekNext() == "*" {
					toMatch++
				}
				if s.peek() == "\n" {
					s.line++
				}
				if toMatch == 0 {
					break
				}
				s.advance()
			}
			if s.peek() == "*" && s.peekNext() == "/" {
				s.advance()
				s.advance()
			} else {
				utils.Check(&utils.LoxError{
					Line: s.line,
					Msg:  "No closing of block comment",
				})
			}
		} else {
			s._addToken(SLASH)
		}
	case " ":
	case "\r":
	case "\t":
		break
	case "\n":
		s.line++
	case "\"":
		s.scanString()
	default:
		if isDigit(c) {
			s.scanNumber()
		} else if isAlpha(c) {
			s.scanIdetifier()
		} else {
			utils.Check(&utils.LoxError{
				Line: s.line,
				Msg:  "Unexpected character '" + string(s.source[s.current]) + "'",
			})
		}
	}
}

func isDigit(c string) bool {
	// change to regex
	return c >= "0" && c <= "9"
}

func isAlpha(c string) bool {
	// change to regex
	return (c >= "a" && c <= "z") || (c >= "A" && c >= "Z") || c == "_"
}

func isAplhNumeric(c string) bool {
	return isAlpha(c) || isDigit(c)
}

func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanIdetifier() {
	for isAplhNumeric(s.peek()) {
		s.advance()
	}
	text := string(s.source[s.start:s.current])
	tok_type, ok := s.Keywords[text]
	if !ok {
		tok_type = IDENTIFIER
	}
	s._addToken(tok_type)
}

func (s *Scanner) scanNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == "." && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	value, err := strconv.ParseFloat(string(s.source[s.start:s.current]), 64)
	utils.Check(err)

	s.addToken(NUMBER, value)
}

func (s *Scanner) scanString() {
	for s.peek() != "\"" && !s.isAtEnd() {
		if s.peek() == "\n" {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		utils.Check(&utils.LoxError{
			Line: s.line,
			Msg:  "Unterminated string",
		})
	}

	s.advance()

	value := string(s.source[s.start+1 : s.current-1])
	s.addToken(STRING, value)
}

func (s *Scanner) peekNext() string {
	// two char lookahead

	if s.current+1 >= len(s.source) {
		return "\000"
	}
	return string(s.source[s.current+1])
}

func (s *Scanner) peek() string {
	// one char lookahead
	if s.isAtEnd() {
		return "\000"
	}
	return string(s.source[s.current])
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})
	return s.tokens
}
