package interpreter

import (
	"testing"

	"pawn/ast"
	"github.com/stretchr/testify/assert"
)

func TestEvaluateMemberIndex(t *testing.T) {
	env := NewEnvironment(nil)
	list := PawnList{&[]PawnValue{PawnBoolean{true}, PawnList{&[]PawnValue{}}}}
	env.variables["varForList"] = list

	listAsVar := ast.MemberExpression{
		Member: ast.Name{Name: "varForList"},
	}

	indexInBounds := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{1}}}
	indexOutOfBounds := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{3}}}

	assert.Equal(t, MemberReturnedIndex{*env.variables["varForList"].(PawnList).v, 1}, evaluateMemberIndex(indexInBounds, env))
	assert.PanicsWithValue(t,
		PawnError{"Cannot index with the integer 3 because it is not less than the list's size 2"},
		func () {evaluateMemberIndex(indexOutOfBounds, env)},
	)
	assert.PanicsWithValue(t,
		PawnError{"Cannot index with the integer 3 because it is not less than the list's size 2"},
		func () {evaluateMemberIndex(indexOutOfBounds, env)},
	)
	indexNonList := ast.Index{ast.BooleanLiteral{true}, []ast.Expression{ast.IntLiteral{3}}}
	assert.PanicsWithValue(t,
		PawnError{"Tried indexing non-list: {true}"},
		func () {evaluateMemberIndex(indexNonList, env)},
	)
}

func TestEvaluateMemberListIndex(t *testing.T) {
	env := NewEnvironment(nil)
	list := PawnList{&[]PawnValue{PawnBoolean{true}, PawnList{}}}
	(*list.v)[1] = list // Recursive list
	env.variables["varForList"] = list
	listAsVar := ast.MemberExpression{
		Member: ast.Name{Name: "varForList"},
	}

	// 0-length empty indexing
	memberIndexListEmpty := ast.Index{listAsVar, []ast.Expression{}}
	assert.Equal(t, MemberReturnedList{PawnList{&[]PawnValue{}}}, evaluateMemberIndex(memberIndexListEmpty, env))

	// duplicate index
	memberIndexListDuplicate := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{0}, ast.IntLiteral{0}}}
	assert.Equal(t, MemberReturnedList{PawnList{&[]PawnValue{PawnBoolean{true}, PawnBoolean{true}}}}, evaluateMemberIndex(memberIndexListDuplicate, env))

	// 3-value indexing
	memberIndexList3Values := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{1}, ast.IntLiteral{0}, ast.IntLiteral{1}}}
	assert.Equal(t, MemberReturnedList{PawnList{&[]PawnValue{list, PawnBoolean{true}, list}}}, evaluateMemberIndex(memberIndexList3Values, env))

	// Index not found
	memberIndexListNotFound := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{0}, ast.IntLiteral{2}}}
	assert.PanicsWithValue(t, PawnError{"Cannot index with the integer 2 because it is not less than the list's size 2"}, func () {
		evaluateMemberIndex(memberIndexListNotFound, env)
	})

	// Multi-indexing non-list
	indexNonList := ast.Index{ast.BooleanLiteral{true}, []ast.Expression{ast.IntLiteral{1}, ast.IntLiteral{1}}}
	assert.PanicsWithValue(t,
		PawnError{"Tried indexing non-list: {true}"},
		func () {evaluateMemberIndex(indexNonList, env)},
	)

	// TODO After implementing call completely, test changing list in second arg while indexing.
	// It's supposed to index as the last step, so index errors do not happen if the non-modified list is short.
}

func TestEvaluateMemberProperty(t *testing.T) {
	env := NewEnvironment(nil)
	object := PawnObject{map[string]PawnValue{"property": PawnBoolean{true}}}
	env.variables["varForObject"] = object
	objectAsVar := ast.MemberExpression{
		Member: ast.Name{Name: "varForObject"},
	}
	assert.Equal(t, MemberReturnedMap{object.v, "property"}, evaluateMemberProperty(ast.Property{Object: objectAsVar, Property: "property"}, env))
	assert.PanicsWithValue(t, PawnError{"Field \"notProperty\" doesn't exist in object"}, func() {
		evaluateMemberProperty(ast.Property{Object: objectAsVar, Property: "notProperty"}, env)
	})

	accesssNonObject := ast.Property{ast.BooleanLiteral{true}, "property"}
	assert.PanicsWithValue(t,
		PawnError{"Tried accessing property of non-object: {map[]}"},
		func () {evaluateMemberProperty(accesssNonObject, env)},
	)
}

func TestEvaluateMemberEnvironmentName(t *testing.T) {
	memberNameFound := ast.Name{Name: "var"}
	memberNameNotFound := ast.Name{Name: "notVar"}

	env := NewEnvironment(nil)
	env.variables["var"] = PawnBoolean{true}
	envDeep := NewEnvironment(&env)

	// Var doesn't exist, shallow and deep search:
	assert.PanicsWithValue(t, PawnError{"Variable \"notVar\" doesn't exist"}, func() {
		evaluateMemberName(memberNameNotFound, env)
	})
	assert.PanicsWithValue(t, PawnError{"Variable \"notVar\" doesn't exist"}, func() {
		evaluateMemberName(memberNameNotFound, envDeep)
	})

	// Var exists, shallow and deep search:
	assert.Equal(t, MemberReturnedMap{env.variables, "var"}, evaluateMemberName(memberNameFound, env))
	assert.Equal(t, MemberReturnedMap{env.variables, "var"}, evaluateMemberName(memberNameFound, envDeep))
}
