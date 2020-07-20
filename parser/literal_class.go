package parser

import (
	"ghostlang.org/ghost/ast"
	"ghostlang.org/ghost/token"
)

func (p *Parser) parseClassLiteral() ast.Expression {
	literal := &ast.ClassLiteral{Token: p.currentToken}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	literal.Name = p.currentToken.Literal

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	literal.Methods = p.parseClassMethods()

	if !p.currentTokenIs(token.RBRACE) {
		return nil
	}

	return literal
}

func (p *Parser) parseClassMethods() []ast.Expression {
	methods := []ast.Expression{}

	for !p.currentTokenIs(token.RBRACE) && !p.currentTokenIs(token.EOF) {
		if p.currentTokenIs(token.FUNCTION) {
			method := p.parseFunctionLiteral()
			methods = append(methods, method)
		}

		p.nextToken()
	}

	return methods
}
