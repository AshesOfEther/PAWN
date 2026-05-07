package interpreter

import (
	"pawn/ast"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeclareUserFunction(t *testing.T) {
	positionalArguments := []string{"x", "y", "z"}
	body := []ast.Statement{
		ast.Assignment{
			ast.Name{"variable"},
			ast.IntLiteral{3},
		},
		ast.Return{nil},
	}
	
	environment := NewEnvironment(nil)
	evaluateFunctionDeclaration(ast.FunctionDecl{
		"foo",
		positionalArguments,
		[]ast.NamedArg{
			ast.NamedArg{"a", ast.IntLiteral{2}},
			ast.NamedArg{"b", ast.FloatLiteral{3}},
		},
		body,
	}, environment)

	assert.Equal(t, PawnFunctionUser{
		&PawnFunctionUserInner{
			environment,
			positionalArguments,
			map[string]ast.Expression{
				"a": ast.IntLiteral{2},
				"b": ast.FloatLiteral{3},
			},
			body,
		},
	}, environment.variables["foo"])
}

func TestDeclareUserFunctionWithDuplicateNamedArgument(t *testing.T) {
	assert.PanicsWithValue(t, duplicateNamedArgumentDeclarationError("foo"), func() {
		evaluateFunctionDeclaration(ast.FunctionDecl{
			"name",
			[]string{},
			[]ast.NamedArg{
				ast.NamedArg{"foo", ast.IntLiteral{1}},
				ast.NamedArg{"bar", ast.IntLiteral{2}},
				ast.NamedArg{"foo", ast.IntLiteral{3}},
			},
			[]ast.Statement{},
		}, NewEnvironment(nil))
	})
}

func TestWhile(t *testing.T) {
	dummyEnv := NewEnvironment(nil)
	whileAST := ast.While{
		Condition: ast.BooleanLiteral{Value: false},
		Body: []ast.Statement{},
	}

	done := false
	go func() {
		evaluateWhile(whileAST, dummyEnv)
		done = true
	}()
	time.Sleep(time.Second / 20)
	assert.True(t, done)
}
