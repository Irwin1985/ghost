package ghost

import (
	"ghostlang.org/x/ghost/builtins"
	"ghostlang.org/x/ghost/evaluator"
	"ghostlang.org/x/ghost/lexer"
	"ghostlang.org/x/ghost/object"
	"ghostlang.org/x/ghost/parser"
	"ghostlang.org/x/ghost/utilities"
	"ghostlang.org/x/ghost/value"
)

var script = Script{}

// Script stores a new Ghost script source.
type Script struct {
	source string
}

// Global values
var (
	// NULL represents a null value.
	NULL = value.NULL

	// TRUE represents a true value.
	TRUE = value.TRUE

	// FALSE represents a false value.
	FALSE = value.FALSE
)

// NewScript registers a new Ghost script to be evaluated.
func NewScript(source string) {
	script = Script{source: source}
}

// RegisterFunction registers a new native function with Ghost.
func RegisterFunction(name string, function object.BuiltinFunction) {
	builtins.RegisterFunction(name, function)
}

// Evaluate runs the registered script through the Ghost evaluator.
func Evaluate() {
	env := object.NewEnvironment()
	l := lexer.New(script.source)
	p := parser.New(l)
	program := p.ParseProgram()

	evaluator.Eval(program, env)
}

// NewError returns a new error object used during runtime.
func NewError(format string, a ...interface{}) *object.Error {
	return utilities.NewError(format, a)
}
