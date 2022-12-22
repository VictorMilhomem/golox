package ast

type LoxFunction struct {
	declaration Function[Types]
}

func NewLoxFunction(decl Function[Types]) LoxFunction {
	return LoxFunction{
		declaration: decl,
	}
}

func (fn LoxFunction) Call(interpreter *Interpreter, arguments []interface{}) (retvalue interface{}) {
	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(ReturnT); ok {
				retvalue = v.Value
				return
			}
			panic(err)
		}
	}()
	env := NewEnvironmentEnclosing(interpreter.globals)

	for i, v := range fn.declaration.Params {
		env.Define(v.Lexeme, arguments[i])
	}
	interpreter.executeBlock(fn.declaration.Body, env)
	return nil
}

func (fn LoxFunction) Arity() int {
	return len(fn.declaration.Params)
}
