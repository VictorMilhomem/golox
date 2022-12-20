package ast

import (
	"github.com/VictorMilhomem/glox/glox/lexer"
	"github.com/VictorMilhomem/glox/glox/utils"
)

type Environment struct {
	enclosing *Environment
	Values    map[string]Types
}

func NewEnvironment() *Environment {
	return &Environment{
		enclosing: nil,
		Values:    make(map[string]Types),
	}
}

func NewEnvironmentEnclosing(enclose *Environment) *Environment {
	return &Environment{
		enclosing: enclose,
		Values:    make(map[string]Types),
	}
}

func (e *Environment) Define(name string, value Types) {
	e.Values[name] = value
}

func (e *Environment) Get(name lexer.Token) Types {
	value, ok := e.Values[name.Lexeme]
	if ok {
		return value
	}
	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}
	utils.Check(NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"' "))
	return nil
}

func (e *Environment) Assign(name lexer.Token, value Types) {
	_, ok := e.Values[name.Lexeme]
	if ok {
		e.Values[name.Lexeme] = value
		return
	}

	if e.enclosing != nil {
		e.enclosing.Assign(name, value)
		return
	}

	utils.Check(NewRuntimeError(name, "Undefined variable '"+name.Lexeme+"'"))
}
