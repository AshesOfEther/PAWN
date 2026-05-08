package interpreter

import (
	"testing"

	"pawn/ast"
	"github.com/stretchr/testify/assert"
)

//func TestEvaluateMemberIndex(t *testing.T) {
//	memberStmt :=	ast.Index{}
//	EvaluateMember(memberStmt, NewEnvironment(nil))
//}
//
//func TestEvaluateMemberListIndex(t *testing.T) {
//	memberStmt :=	ast.Index{}
//	EvaluateMember(memberStmt, NewEnvironment(nil))
//}
//
//func TestEvaluateMemberProperty(t *testing.T) {
//	memberStmt :=	ast.Property{}
//	EvaluateMember(memberStmt, NewEnvironment(nil))
//}

func TestEvaluateMemberEnvironmentName(t *testing.T) {
	memberStmtFound := ast.Name{Name: "var"}
	memberStmtNotFound := ast.Name{Name: "notVar"}

	env := NewEnvironment(nil)
	env.variables["var"] = PawnBoolean{true}
	envDeep := NewEnvironment(&env)
	_ = envDeep
	_ = memberStmtNotFound

	// Var doesn't exist, shallow and deep search:
	assert.PanicsWithValue(t, PawnError{"Variable \"notVar\" doesn't exist"}, func() {
		EvaluateMember(memberStmtNotFound, env)
	})
	assert.PanicsWithValue(t, PawnError{"Variable \"notVar\" doesn't exist"}, func() {
		EvaluateMember(memberStmtNotFound, envDeep)
	})

	// Var exists, shallow and deep search:
	assert.Equal(t, MemberReturnedMap{env.variables, "var"}, EvaluateMember(memberStmtFound, env))
	assert.Equal(t, MemberReturnedMap{env.variables, "var"}, EvaluateMember(memberStmtFound, envDeep))
}
