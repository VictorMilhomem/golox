package main

import (
	"fmt"
	"strings"

	"github.com/VictorMilhomem/glox/ast"
	"github.com/VictorMilhomem/glox/glox/lexer"
)

type AstPrinter struct{}

func (a *AstPrinter) print(expr ast.IExpr[string]) string {
	return fmt.Sprintf("%v", expr.Accept(a))
}

func (a *AstPrinter) VisitBinary(expr ast.Binary[string]) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func NewBinary(left ast.IExpr[string], right ast.IExpr[string], token *lexer.Token) *ast.Binary[string] {
	return &ast.Binary[string]{
		Left:     left,
		Operator: *token,
		Right:    right,
	}
}

func (a *AstPrinter) VisitUnary(expr ast.Unary[string]) interface{} {
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func NewUnary(op *lexer.Token, right ast.IExpr[string]) *ast.Unary[string] {
	return &ast.Unary[string]{
		Operator: *op,
		Right:    right,
	}
}

func (a *AstPrinter) VisitGrouping(expr ast.Grouping[string]) interface{} {
	return a.parenthesize("group", expr.Expression)
}

func NewGrouping(expr ast.IExpr[string]) *ast.Grouping[string] {
	return &ast.Grouping[string]{
		Expression: expr,
	}
}

func (a *AstPrinter) VisitLiteral(expr ast.Literal[string]) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprint(expr.Value)
}

func NewLiteral(value lexer.Object) *ast.Literal[string] {
	return &ast.Literal[string]{
		Value: value,
	}
}

func (a *AstPrinter) parenthesize(name string, exprs ...ast.IExpr[string]) string {
	builder := strings.Builder{}

	builder.WriteString("(" + name)
	for _, e := range exprs {
		builder.WriteString(" ")
		builder.WriteString(a.print(e))
	}
	builder.WriteString(")")
	return builder.String()
}

func main() {
	printer := &AstPrinter{}
	l := NewLiteral("1234")
	unary := NewUnary(lexer.NewToken(lexer.MINUS, "-", nil, 1), l)
	star := lexer.NewToken(lexer.STAR, "*", nil, 1)
	l2 := NewLiteral("12")
	group := NewGrouping(l2)

	binary := NewBinary(unary, group, star)
	fmt.Println(printer.print(binary))
}
