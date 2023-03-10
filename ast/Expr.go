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
type Assign[T Types] struct {
    Name Token
    Value Expr[T]
}
func (v *Assign[T])Accept(visitor ExprVisitor[T]) interface{} {
    return visitor.VisitAssign(*v)
}
func NewAssign(name Token,value Expr[Types],) *Assign[Types]{
    return &Assign[Types]{
    Name: name,
    Value: value,
    }
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
type Call[T Types] struct {
    Callee Expr[T]
    Paren Token
    Arguments []Expr[Types]
}
func (v *Call[T])Accept(visitor ExprVisitor[T]) interface{} {
    return visitor.VisitCall(*v)
}
func NewCall(callee Expr[Types],paren Token,arguments []Expr[Types],) *Call[Types]{
    return &Call[Types]{
    Callee: callee,
    Paren: paren,
    Arguments: arguments,
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
type Logical[T Types] struct {
    Left Expr[T]
    Operator Token
    Right Expr[T]
}
func (v *Logical[T])Accept(visitor ExprVisitor[T]) interface{} {
    return visitor.VisitLogical(*v)
}
func NewLogical(left Expr[Types],operator Token,right Expr[Types],) *Logical[Types]{
    return &Logical[Types]{
    Left: left,
    Operator: operator,
    Right: right,
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
type Variable[T Types] struct {
    Name Token
}
func (v *Variable[T])Accept(visitor ExprVisitor[T]) interface{} {
    return visitor.VisitVariable(*v)
}
func NewVariable(name Token,) *Variable[Types]{
    return &Variable[Types]{
    Name: name,
    }
}
type ExprVisitor[T Types] interface{
    VisitAssign(expr Assign[T]) interface{}
    VisitBinary(expr Binary[T]) interface{}
    VisitCall(expr Call[T]) interface{}
    VisitGrouping(expr Grouping[T]) interface{}
    VisitLiteral(expr Literal[T]) interface{}
    VisitLogical(expr Logical[T]) interface{}
    VisitUnary(expr Unary[T]) interface{}
    VisitVariable(expr Variable[T]) interface{}
}
