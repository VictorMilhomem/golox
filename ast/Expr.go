package ast

import (
	. "github.com/VictorMilhomem/glox/glox/lexer"
	"golang.org/x/exp/constraints"
)

type Types interface {
	constraints.Ordered
}
type IExpr[T Types] interface {
	Accept(visitor ExprVisitor[T]) interface{}
}
type Binary[T Types] struct {
	Left     IExpr[T]
	Operator Token
	Right    IExpr[T]
}

func (v *Binary[T]) Accept(visitor ExprVisitor[T]) interface{} {
	return visitor.VisitBinary(*v)
}

type Grouping[T Types] struct {
	Expression IExpr[T]
}

func (v *Grouping[T]) Accept(visitor ExprVisitor[T]) interface{} {
	return visitor.VisitGrouping(*v)
}

type Literal[T Types] struct {
	Value Object
}

func (v *Literal[T]) Accept(visitor ExprVisitor[T]) interface{} {
	return visitor.VisitLiteral(*v)
}

type Unary[T Types] struct {
	Operator Token
	Right    IExpr[T]
}

func (v *Unary[T]) Accept(visitor ExprVisitor[T]) interface{} {
	return visitor.VisitUnary(*v)
}

type ExprVisitor[T Types] interface {
	VisitBinary(expr Binary[T]) interface{}
	VisitGrouping(expr Grouping[T]) interface{}
	VisitLiteral(expr Literal[T]) interface{}
	VisitUnary(expr Unary[T]) interface{}
}
