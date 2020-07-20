package object

import (
	"bytes"

	"ghostlang.org/ghost/ast"
)

type Class struct {
	Name    string
	Methods []*ast.Expression
	Env     *Environment
}

func (c *Class) Type() ObjectType {
	return CLASS_OBJ
}

func (c *Class) Inspect() string {
	var out bytes.Buffer

	out.WriteString("class\n")

	return out.String()
}
