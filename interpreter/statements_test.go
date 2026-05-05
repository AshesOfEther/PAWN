package interpreter

import (
	"pawn/ast"
	"testing"
)

func TestWhile(t *testing.T) {
	dummyEnv := NewEnvironment(nil)
	whileAST := ast.While{
		Condition: ast.BooleanLiteral{Value: false},
		Body: []ast.Statement{},
	}

	evaluateWhile(whileAST, dummyEnv)
}
