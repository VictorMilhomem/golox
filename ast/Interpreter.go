package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VictorMilhomem/glox/glox/lexer"
	"github.com/VictorMilhomem/glox/glox/utils"
)

type Interpreter struct {
	env *Environment
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		env: NewEnvironment(),
	}
}

func (i *Interpreter) Interpret(statements []Stmt[Types]) {
	for _, stmt := range statements {
		i.execute(stmt)
	}
}

func (i *Interpreter) execute(stmt Stmt[Types]) {
	stmt.Accept(i)
}

/*
* VISIT EXPRESSIONS
 */

func (i *Interpreter) VisitLiteral(expr Literal[Types]) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitGrouping(expr Grouping[Types]) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnary(expr Unary[Types]) interface{} {
	right := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case lexer.BANG:
		return !i.isTruly(right)
	case lexer.MINUS:
		i.checkNumberOperand(expr.Operator, right)
		return -float64(right.(float64))
	}
	return nil
}

func (i *Interpreter) VisitAssign(expr Assign[Types]) interface{} {
	value := i.evaluate(expr.Value)
	i.env.Assign(expr.Name, value)
	return value
}

func (i *Interpreter) VisitBinary(expr Binary[Types]) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case lexer.GREATER:
		i.checkNumberOperands(expr.Operator, left, right)
		return float64(left.(float64)) > float64(right.(float64))
	case lexer.GREATER_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return float64(left.(float64)) >= float64(right.(float64))
	case lexer.LESS:
		i.checkNumberOperands(expr.Operator, left, right)
		return float64(left.(float64)) < float64(right.(float64))
	case lexer.LESS_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return float64(left.(float64)) <= float64(right.(float64))
	case lexer.BANG_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return !i.isEqual(left, right)
	case lexer.EQUAL_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return i.isEqual(left, right)
	case lexer.MINUS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) - right.(float64)
	case lexer.PLUS:
		l, okl := left.(string)
		r, okr := right.(string)
		if okl && okr {
			return string(l) + string(r)
		}

		lf, okfl := left.(float64)
		rr, okfr := right.(float64)
		if okfl && okfr {
			return float64(lf) + float64(rr)
		}
		utils.Check(NewRuntimeError(expr.Operator, "Operands must be two numbers or two strings."))
	case lexer.SLASH:
		i.checkNumberOperands(expr.Operator, left, right)
		return float64(left.(float64)) / float64(right.(float64))
	case lexer.STAR:
		i.checkNumberOperands(expr.Operator, left, right)
		return float64(left.(float64)) * float64(right.(float64))
	}

	return nil
}

func (i *Interpreter) VisitVariable(expr Variable[Types]) interface{} {
	return i.env.Get(expr.Name)
}

/*
* VISIT STATEMENTS
 */

func (i *Interpreter) VisitBlock(stmt Block[Types]) interface{} {
	i.executeBlock(stmt.Statements, NewEnvironmentEnclosing(i.env))
	return nil
}

func (i *Interpreter) executeBlock(stmts []Stmt[Types], env *Environment) {
	previous := i.env
	i.env = env
	for _, statement := range stmts {
		i.execute(statement)
	}
	i.env = previous
}

func (i *Interpreter) VisitExpression(stmt Expression[Types]) interface{} {
	i.evaluate(stmt.Expression)
	return nil
}

func (i *Interpreter) VisitPrint(stmt Print[Types]) interface{} {
	value := i.evaluate(stmt.Expression)
	fmt.Println(i.stringify(value))
	return nil
}

func (i *Interpreter) VisitVar(stmt Var[Types]) interface{} {
	var value Types
	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}
	i.env.Define(stmt.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) evaluate(expr Expr[Types]) Types {
	return expr.Accept(i)
}

func (i *Interpreter) isEqual(a Types, b Types) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}

	return true
}

func (i *Interpreter) isTruly(obj Types) bool {
	if obj == nil {
		return false
	}
	v, ok := obj.(bool)
	if ok {
		return bool(v)
	}
	return true
}

func (i *Interpreter) checkNumberOperand(operator lexer.Token, operand Types) {
	_, ok := operand.(float64)
	if ok {
		return
	}
	utils.Check(NewParserError(operator, "Operand must be a number"))
}

func (i *Interpreter) checkNumberOperands(operator lexer.Token, left Types, right Types) {
	_, okl := left.(float64)
	_, okr := right.(float64)

	if okl && okr {
		return
	}

	utils.Check(NewRuntimeError(operator, "Operands must be numbers"))
}

func (i *Interpreter) stringify(obj Types) string {
	if obj == nil {
		return "nil"
	}
	v, ok := obj.(float64)
	if ok {
		precision := 4
		text := strconv.FormatFloat(v, 'f', precision, 64)

		if strings.Contains(text, ".00") {
			text = strings.Split(text, ".")[0]
		}
		return text
	}
	return fmt.Sprintf("%v", obj)
}
