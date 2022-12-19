package ast

import (
	"github.com/VictorMilhomem/glox/glox/lexer"
	"github.com/VictorMilhomem/glox/glox/utils"
)

type Parser struct {
	tokens  []lexer.Token
	current int
}

func NewParser(t []lexer.Token) *Parser {
	return &Parser{
		tokens:  t,
		current: 0,
	}
}

func (p *Parser) Parse() Expr[Types] {
	expr := p.expression()
	if expr != nil {
		return expr
	}
	return nil
}

func (p *Parser) expression() Expr[Types] {
	return p.equality()
}

func (p *Parser) equality() Expr[Types] {
	expr := p.comparison()

	for p.match(lexer.BANG, lexer.BANG_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) comparison() Expr[Types] {
	expr := p.term()
	for p.match(lexer.GREATER, lexer.GREATER_EQUAL, lexer.LESS, lexer.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) term() Expr[Types] {
	expr := p.factor()
	for p.match(lexer.MINUS, lexer.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) factor() Expr[Types] {
	expr := p.unary()
	for p.match(lexer.SLASH, lexer.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = NewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) unary() Expr[Types] {
	if p.match(lexer.BANG, lexer.MINUS) {
		operator := p.previous()
		right := p.unary()
		expr := NewUnary(operator, right)
		return expr
	}
	return p.primary()
}

func (p *Parser) primary() Expr[Types] {
	if p.match(lexer.FALSE) {
		return NewLiteral(false)
	}
	if p.match(lexer.TRUE) {
		return NewLiteral(true)
	}
	if p.match(lexer.NIL) {
		return NewLiteral(nil)
	}

	if p.match(lexer.NUMBER, lexer.STRING) {
		return NewLiteral(p.previous().Literal)
	}

	if p.match(lexer.LEFT_PAREN) {
		expr := p.expression()
		p.consume(lexer.RIGHT_PAREN, "Expected ')' after expression")
		return NewGrouping(expr)
	}
	utils.Check(NewParserError(p.peek(), "Expected expression"))
	return nil
}

func (p *Parser) consume(tp lexer.TokenType, msg string) lexer.Token {
	if !p.check(tp) {
		utils.Check(NewParserError(p.peek(), msg))
	}

	return p.advance()
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == lexer.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case lexer.CLASS:
		case lexer.FUN:
		case lexer.FOR:
		case lexer.VAR:
		case lexer.IF:
		case lexer.WHILE:
		case lexer.PRINT:
		case lexer.RETURN:
			return
		}
		p.advance()
	}
}

func (p *Parser) match(types ...lexer.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(t lexer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() lexer.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == lexer.EOF
}

func (p *Parser) peek() lexer.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() lexer.Token {
	return p.tokens[p.current-1]
}
