package ast
import(
. "github.com/VictorMilhomem/glox/glox/lexer"
"golang.org/x/exp/constraints"
)
type Node interface{
    constraints.Ordered
}
type Expr[T Node] interface {
    Accept(visitor ExprVisitor[T]) T
}
type Binary[T Node] struct {
    Left *Expr[T]
    Operator *Token
    Right *Expr[T]
}
func (v *Binary[T])Accept(visitor ExprVisitor[T]) T {
    return visitor.visitBinary(v)
}
type Grouping[T Node] struct {
    Expression *Expr[T]
}
func (v *Grouping[T])Accept(visitor ExprVisitor[T]) T {
    return visitor.visitGrouping(v)
}
type Literal[T Node] struct {
    Value *Object
}
func (v *Literal[T])Accept(visitor ExprVisitor[T]) T {
    return visitor.visitLiteral(v)
}
type Unary[T Node] struct {
    Operator *Token
    Right *Expr[T]
}
func (v *Unary[T])Accept(visitor ExprVisitor[T]) T {
    return visitor.visitUnary(v)
}
type ExprVisitor[T Node] interface{
    visitBinary(expr *Binary[T])T
    visitGrouping(expr *Grouping[T])T
    visitLiteral(expr *Literal[T])T
    visitUnary(expr *Unary[T])T
}
