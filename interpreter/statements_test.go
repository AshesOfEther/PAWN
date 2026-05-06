package interpreter

import (
	"pawn/ast"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
