package lexer

import (
	"fmt"
)

type TokenType int16

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF = -1
)

func (t *TokenType) toString() string {
	var name string
	switch *t {
	case LEFT_PAREN:
		name = "LEFT_PAREN"
	case RIGHT_PAREN:
		name = "RIGHT_PAREN"
	case LEFT_BRACE:
		name = "LEFT_BRACE"
	case RIGHT_BRACE:
		name = "RIGHT_PAREN"
	case COMMA:
		name = "COMMA"
	case DOT:
		name = "DOT"
	case MINUS:
		name = "MINUS"
	case PLUS:
		name = "PLUS"
	case SEMICOLON:
		name = "SEMICOLON"
	case SLASH:
		name = "SLASH"
	case STAR:
		name = "STAR"
	case BANG:
		name = "BANG"
	case BANG_EQUAL:
		name = "BANG_EQUAL"
	case EQUAL:
		name = "EQUAL"
	case EQUAL_EQUAL:
		name = "EQUAL_EQUAL"
	case GREATER:
		name = "GREATER"
	case GREATER_EQUAL:
		name = "GREATER_EQUAL"
	case LESS:
		name = "LESS"
	case LESS_EQUAL:
		name = "LESS_EQUAL"
	case IDENTIFIER:
		name = "ID"
	case STRING:
		name = "STRING"
	case NUMBER:
		name = "NUMBER"
	case AND:
		name = "AND"
	case CLASS:
		name = "CLASS"
	case ELSE:
		name = "ELSE"
	case FALSE:
		name = "FALSE"
	case FUN:
		name = "FUN"
	case FOR:
		name = "FOR"
	case IF:
		name = "IF"
	case NIL:
		name = "NIL"
	case OR:
		name = "OR"
	case PRINT:
		name = "PRINT"
	case RETURN:
		name = "RETURN"
	case SUPER:
		name = "SUPER"
	case THIS:
		name = "THIS"
	case TRUE:
		name = "TRUE"
	case VAR:
		name = "VAR"
	case WHILE:
		name = "WHILE"
	case EOF:
		name = "EOF"
	}
	return name
}

type Object interface{}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal Object
	Line    int
}

func NewToken(t TokenType, lexeme string, literal string, line int) *Token {
	return &Token{
		Type:    t,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

func (t *Token) ToString() string {
	return fmt.Sprintf("{<Type, Lexeme ,Literal>: <%v, %v, %v>}", t.Type.toString(), t.Lexeme, t.Literal)
}
