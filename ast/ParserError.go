package ast

import (
	"strconv"

	"github.com/VictorMilhomem/glox/glox/lexer"
)

type ParserError struct {
	token lexer.Token
	msg   string
}

type RuntimeError struct {
	token lexer.Token
	msg   string
}

func (e *ParserError) Error() string {
	if e.token.Type == lexer.EOF {
		return strconv.Itoa(e.token.Line) + " at end " + e.msg
	}
	return "Error at Line: " + strconv.Itoa(e.token.Line) + " at '" + e.token.Lexeme + "' " + e.msg
}

func NewParserError(token lexer.Token, msg string) *ParserError {
	return &ParserError{
		token: token,
		msg:   msg,
	}
}

func (e *RuntimeError) Error() string {
	return "RuntimeError at Line: " + strconv.Itoa(e.token.Line) + " '" + e.token.Lexeme + "' " + e.msg
}

func NewRuntimeError(token lexer.Token, msg string) *RuntimeError {
	return &RuntimeError{
		token: token,
		msg:   msg,
	}
}
