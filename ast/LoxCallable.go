package ast

type LoxCallable interface {
	Arity() int
	Call(interpreter *Interpreter, arguments []interface{}) interface{}
}
