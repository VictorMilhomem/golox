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
type Function[T Types] struct {
    Name Token
    Params []Token
    Body []Stmt[Types]
}
func (v *Function[T])Accept(visitor StmtVisitor[T]) interface{} {
    return visitor.VisitFunction(*v)
}
func NewFunction(name Token,params []Token,body []Stmt[Types],) *Function[Types]{
    return &Function[Types]{
    Name: name,
    Params: params,
    Body: body,
    }
}
type If[T Types] struct {
    Condition Expr[T]
    Thenbranch Stmt[T]
    Elsebranch Stmt[T]
}
func (v *If[T])Accept(visitor StmtVisitor[T]) interface{} {
    return visitor.VisitIf(*v)
}
func NewIf(condition Expr[Types],thenBranch Stmt[Types],elseBranch Stmt[Types],) *If[Types]{
    return &If[Types]{
    Condition: condition,
    Thenbranch: thenBranch,
    Elsebranch: elseBranch,
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
type Return[T Types] struct {
    Keyword Token
    Value Expr[T]
}
func (v *Return[T])Accept(visitor StmtVisitor[T]) interface{} {
    return visitor.VisitReturn(*v)
}
func NewReturn(keyword Token,value Expr[Types],) *Return[Types]{
    return &Return[Types]{
    Keyword: keyword,
    Value: value,
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
type While[T Types] struct {
    Condition Expr[T]
    Body Stmt[T]
}
func (v *While[T])Accept(visitor StmtVisitor[T]) interface{} {
    return visitor.VisitWhile(*v)
}
func NewWhile(condition Expr[Types],body Stmt[Types],) *While[Types]{
    return &While[Types]{
    Condition: condition,
    Body: body,
    }
}
type StmtVisitor[T Types] interface{
    VisitBlock(stmt Block[T]) interface{}
    VisitExpression(stmt Expression[T]) interface{}
    VisitFunction(stmt Function[T]) interface{}
    VisitIf(stmt If[T]) interface{}
    VisitPrint(stmt Print[T]) interface{}
    VisitReturn(stmt Return[T]) interface{}
    VisitVar(stmt Var[T]) interface{}
    VisitWhile(stmt While[T]) interface{}
}
