package ast
import(
. "github.com/VictorMilhomem/glox/glox/lexer"
"golang.org/x/exp/constraints"
)
type Types interface{
    constraints.Ordered | Object
}
type Expr[T Types] interface {
    Accept(visitor ExprVisitor[T])  interface{}
}
type Binary[T Types] struct {
    Left Expr[T]
    Operator Token
    Right Expr[T]
}
func (v *Binary[T])Accept(visitor ExprVisitor[T]) interface{} {
    return visitor.VisitBinary(*v)
}
func NewBinary(left Expr[Types],operator Token,right Expr[Types],) *Binary[Types]{
    return &Binary[Types]{
    Left: left,
    Operator: operator,
    Right: right,
    }
}
type Grouping[T Types] struct {
    Expression Expr[T]
}
func (v *Grouping[T])Accept(visitor ExprVisitor[T]) interface{} {
    return visitor.VisitGrouping(*v)
}
func NewGrouping(expression Expr[Types],) *Grouping[Types]{
    return &Grouping[Types]{
    Expression: expression,
    }
}
type Literal[T Types] struct {
    Value Object
}
func (v *Literal[T])Accept(visitor ExprVisitor[T]) interface{} {
    return visitor.VisitLiteral(*v)
}
func NewLiteral(value Object,) *Literal[Types]{
    return &Literal[Types]{
    Value: value,
    }
}
type Unary[T Types] struct {
    Operator Token
    Right Expr[T]
}
func (v *Unary[T])Accept(visitor ExprVisitor[T]) interface{} {
    return visitor.VisitUnary(*v)
}
func NewUnary(operator Token,right Expr[Types],) *Unary[Types]{
    return &Unary[Types]{
    Operator: operator,
    Right: right,
    }
}
type ExprVisitor[T Types] interface{
    VisitBinary(expr Binary[T]) interface{}
    VisitGrouping(expr Grouping[T]) interface{}
    VisitLiteral(expr Literal[T]) interface{}
    VisitUnary(expr Unary[T]) interface{}
}
