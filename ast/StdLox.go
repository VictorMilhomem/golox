package ast

import "time"

type StdLib struct {
	env *Environment
}

func NewStdLib() *StdLib {
	return &StdLib{
		env: NewEnvironment(),
	}
}

func (std *StdLib) Globals() *Environment {
	std.env.Define("clock", Clock{})
	return std.env
}

type Clock struct{}

// Arity() int
// Call(interpreter *Interpreter, arguments []Types) Types

func (clock *Clock) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	return float64(time.Now().UnixMilli())
}

func (clock *Clock) Arity() int {
	return 0
}
