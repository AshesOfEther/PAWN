package interpreter

import (
	"testing"

	"pawn/ast"
	"github.com/stretchr/testify/assert"
)

//func TestEvaluateMemberIndex(t *testing.T) {
//	memberStmt := ast.Index{}
//	EvaluateMember(memberStmt, NewEnvironment(nil))
//}
//
//func TestEvaluateMemberListIndex(t *testing.T) {
//	memberStmt := ast.Index{}
//	EvaluateMember(memberStmt, NewEnvironment(nil))
//}

func TestEvaluateMemberProperty(t *testing.T) {
	env := NewEnvironment(nil)
	object := PawnObject{map[string]PawnValue{"property": PawnBoolean{true}}}
	env.variables["varForObject"] = object
	objectAsVar := ast.MemberExpression{
		Member: ast.Name{Name: "varForObject"},
	}
	assert.Equal(t, MemberReturnedMap{object.v, "property"}, EvaluateMember(ast.Property{Object: objectAsVar, Property: "property"}, env))
	assert.PanicsWithValue(t, PawnError{"Field \"notProperty\" doesn't exist in object"}, func() {
		EvaluateMember(ast.Property{Object: objectAsVar, Property: "notProperty"}, env)
	})
}

func TestEvaluateMemberEnvironmentName(t *testing.T) {
	memberNameFound := ast.Name{Name: "var"}
	memberNameNotFound := ast.Name{Name: "notVar"}

	env := NewEnvironment(nil)
	env.variables["var"] = PawnBoolean{true}
	envDeep := NewEnvironment(&env)

	// Var doesn't exist, shallow and deep search:
	assert.PanicsWithValue(t, PawnError{"Variable \"notVar\" doesn't exist"}, func() {
		EvaluateMember(memberNameNotFound, env)
	})
	assert.PanicsWithValue(t, PawnError{"Variable \"notVar\" doesn't exist"}, func() {
		EvaluateMember(memberNameNotFound, envDeep)
	})

	// Var exists, shallow and deep search:
	assert.Equal(t, MemberReturnedMap{env.variables, "var"}, EvaluateMember(memberNameFound, env))
	assert.Equal(t, MemberReturnedMap{env.variables, "var"}, EvaluateMember(memberNameFound, envDeep))
}
