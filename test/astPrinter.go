package test

import (
	"fmt"
	"strings"

	"github.com/VictorMilhomem/glox/ast"
	"github.com/VictorMilhomem/glox/glox/lexer"
)

type AstPrinter struct{}

func (a *AstPrinter) Print(expr ast.Expr[ast.Types]) ast.Types {
	return fmt.Sprintf("%v", expr.Accept(a))
}

func (a *AstPrinter) VisitBinary(expr ast.Binary[ast.Types]) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func NewBinary(left ast.Expr[ast.Types], right ast.Expr[ast.Types], token *lexer.Token) *ast.Binary[ast.Types] {
	return &ast.Binary[ast.Types]{
		Left:     left,
		Operator: *token,
		Right:    right,
	}
}

func (a *AstPrinter) VisitUnary(expr ast.Unary[ast.Types]) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func NewUnary(op *lexer.Token, right ast.Expr[ast.Types]) *ast.Unary[ast.Types] {
	return &ast.Unary[ast.Types]{
		Operator: *op,
		Right:    right,
	}
}

func (a *AstPrinter) VisitGrouping(expr ast.Grouping[ast.Types]) interface{} {
	return a.parenthesize("group", expr.Expression)
}

func NewGrouping(expr ast.Expr[ast.Types]) *ast.Grouping[ast.Types] {
	return &ast.Grouping[ast.Types]{
		Expression: expr,
	}
}

func (a *AstPrinter) VisitLiteral(expr ast.Literal[ast.Types]) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprint(expr.Value)
}

func NewLiteral(value lexer.Object) *ast.Literal[ast.Types] {
	return &ast.Literal[ast.Types]{
		Value: value,
	}
}

func (a *AstPrinter) parenthesize(name ast.Types, exprs ...ast.Expr[ast.Types]) ast.Types {
	builder := strings.Builder{}

	builder.WriteString("(" + name.(string))
	for _, e := range exprs {
		builder.WriteString(" ")
		builder.WriteString(a.Print(e).(string))
	}
	builder.WriteString(")")
	return builder.String()
}
