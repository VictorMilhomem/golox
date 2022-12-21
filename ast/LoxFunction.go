package ast

type LoxFunction struct {
	declaration Function[Types]
}

func NewLoxFunction(decl Function[Types]) *LoxFunction {
	return &LoxFunction{
		declaration: decl,
	}
}

func (fn *LoxFunction) Call(interpreter *Interpreter, arguments []Types) (retvalue interface{}) {
	env := interpreter.env
	size := len(fn.declaration.Params)

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(ReturnT); ok {
				retvalue = v.Value
				return
			}
			panic(err)
		}
	}()
	for i := 0; i < size; i++ {
		env.Define(fn.declaration.Params[i].Lexeme, arguments[i])
	}
	interpreter.executeBlock(fn.declaration.Body, env)
	return nil
}

func (fn *LoxFunction) Arity() int {
	return len(fn.declaration.Params)
}