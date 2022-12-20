package ast
import(
. "github.com/VictorMilhomem/glox/glox/lexer"
)
type Stmt[T Types] interface {
    Accept(visitor StmtVisitor[T])  interface{}
}
type Block[T Types] struct {
    Statements []Stmt[Types]
}
func (v *Block[T])Accept(visitor StmtVisitor[T]) interface{} {
    return visitor.VisitBlock(*v)
}
func NewBlock(statements []Stmt[Types],) *Block[Types]{
    return &Block[Types]{
    Statements: statements,
    }
}
type Expression[T Types] struct {
    Expression Expr[T]
}
func (v *Expression[T])Accept(visitor StmtVisitor[T]) interface{} {
    return visitor.VisitExpression(*v)
}
func NewExpression(expression Expr[Types],) *Expression[Types]{
    return &Expression[Types]{
    Expression: expression,
    }
}
type Print[T Types] struct {
    Expression Expr[T]
}
func (v *Print[T])Accept(visitor StmtVisitor[T]) interface{} {
    return visitor.VisitPrint(*v)
}
func NewPrint(expression Expr[Types],) *Print[Types]{
    return &Print[Types]{
    Expression: expression,
    }
}
type Var[T Types] struct {
    Name Token
    Initializer Expr[T]
}
func (v *Var[T])Accept(visitor StmtVisitor[T]) interface{} {
    return visitor.VisitVar(*v)
}
func NewVar(name Token,initializer Expr[Types],) *Var[Types]{
    return &Var[Types]{
    Name: name,
    Initializer: initializer,
    }
}
type StmtVisitor[T Types] interface{
    VisitBlock(stmt Block[T]) interface{}
    VisitExpression(stmt Expression[T]) interface{}
    VisitPrint(stmt Print[T]) interface{}
    VisitVar(stmt Var[T]) interface{}
}
