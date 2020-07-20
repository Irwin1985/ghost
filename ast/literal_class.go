package ast

import (
	"bytes"
	"strings"

	"ghostlang.org/ghost/token"
)

type ClassLiteral struct {
	Token   token.Token
	Name    string
	Methods []Expression
}

func (cl *ClassLiteral) expressionNode() {}

func (cl *ClassLiteral) TokenLiteral() string {
	return cl.Token.Literal
}

func (cl *ClassLiteral) String() string {
	var out bytes.Buffer

	methods := []string{}

	for _, p := range cl.Methods {
		methods = append(methods, p.String())
	}

	out.WriteString(cl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(methods, ", "))
	out.WriteString(") ")

	return out.String()
}
