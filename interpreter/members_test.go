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

	nonIndex := ast.Index{listAsVar, []ast.Expression{ast.StringLiteral{"nonIndex"}}}
	indexInBounds := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{1}}}
	indexOutOfBounds := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{3}}}

	assert.PanicsWithValue(t,
		typeError("int", PawnString{"nonIndex"}),
		func () {evaluateMemberIndex(nonIndex, env)},
	)
	assert.Equal(t, Index{*env.variables["varForList"].(PawnList).v, 1}, evaluateMemberIndex(indexInBounds, env))
	assert.PanicsWithValue(t,
		indexTooHighError(3, env.variables["varForList"].(PawnList)),
		func () {evaluateMemberIndex(indexOutOfBounds, env)},
	)
	indexNonList := ast.Index{ast.BooleanLiteral{true}, []ast.Expression{ast.IntLiteral{3}}}
	assert.PanicsWithValue(t,
		typeError("list", PawnBoolean{true}),
		func () {evaluateMemberIndex(indexNonList, env)},
	)
}

func listIndexTestInit() (listAsVar ast.MemberExpression, listInner, list PawnList, env Environment)  {
	env = NewEnvironment(nil)
	list = PawnList{&[]PawnValue{PawnBoolean{true}, PawnList{}}}
	listInner = PawnList{&[]PawnValue{PawnBoolean{true}}}
	(*list.v)[1] = listInner // List with a reference to test that list duplication is not happening
	env.variables["varForList"] = list
	listAsVar = ast.MemberExpression{
		Member: ast.Name{Name: "varForList"},
	}
	return
}

func TestEvaluateMemberListIndexNonIntError(t *testing.T) {
	listAsVar, _, _, env := listIndexTestInit()

	// Non-int indexing
	nonIntIndexing := ast.Index{listAsVar, []ast.Expression{
		ast.IntLiteral{0}, ast.StringLiteral{"nonInt"}},
	}
	assert.PanicsWithValue(t,
		typeError("integer", PawnString{"nonInt"}),
		func () {evaluateMemberIndex(nonIntIndexing, env)},
	)
}

func TestEvaluateMemberListIndexEmptyInvalid(t *testing.T) {
	listAsVar, _, _, env := listIndexTestInit()

	// 0-length empty indexing
	memberIndexListEmpty := ast.Index{listAsVar, []ast.Expression{}}
	assert.Panics(t, func () {evaluateMemberIndex(memberIndexListEmpty, env)})
}

func TestEvaluateMemberListIndexDuplicate(t *testing.T) {
	listAsVar, _, _, env := listIndexTestInit()

	// duplicate index
	memberIndexListDuplicate := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{0}, ast.IntLiteral{0}}}
	assert.Equal(t, PawnList{&[]PawnValue{PawnBoolean{true}, PawnBoolean{true}}}, evaluateMemberIndex(memberIndexListDuplicate, env))
}

func TestEvaluateMemberListIndex3Values(t *testing.T) {
	listAsVar, listInner, _, env := listIndexTestInit()

	// 3-value indexing
	memberIndexList3Values := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{1}, ast.IntLiteral{0}, ast.IntLiteral{1}}}
	assert.Equal(t, PawnList{&[]PawnValue{listInner, PawnBoolean{true}, listInner}}, evaluateMemberIndex(memberIndexList3Values, env))
}

func TestEvaluateMemberListIndexOutOfBounds(t *testing.T) {
	listAsVar, _, list, env := listIndexTestInit()

	// Index not found
	memberIndexListNotFound := ast.Index{listAsVar, []ast.Expression{ast.IntLiteral{0}, ast.IntLiteral{2}}}
	assert.PanicsWithValue(t, indexTooHighError(2, list), func () {
		evaluateMemberIndex(memberIndexListNotFound, env)
	})
}

func TestEvaluateMemberListIndex(t *testing.T) {
	_, _, _, env := listIndexTestInit()

	// Multi-indexing non-list
	indexNonList := ast.Index{ast.BooleanLiteral{true}, []ast.Expression{ast.IntLiteral{1}, ast.IntLiteral{1}}}
	assert.PanicsWithValue(t,
		typeError("list", PawnBoolean{true}),
		func () {evaluateMemberIndex(indexNonList, env)},
	)

	// TODO After implementing function call with body completely, test changing list in second arg while indexing.
	// It's supposed to index as the last step, so index errors do not happen if the non-modified list is short.
}

func TestEvaluateMemberProperty(t *testing.T) {
	env := NewEnvironment(nil)
	object := PawnObject{map[string]PawnValue{"property": PawnBoolean{true}}}
	env.variables["varForObject"] = object
	objectAsVar := ast.MemberExpression{
		Member: ast.Name{Name: "varForObject"},
	}
	assert.Equal(t,
		PropertyOrVar{object.v, "property"},
		evaluateMemberProperty(ast.Property{Object: objectAsVar, Property: "property"}, env),
	)
	assert.PanicsWithValue(t,
		fieldNonExistantError("notProperty"),
		func() {
			evaluateMemberProperty(ast.Property{Object: objectAsVar, Property: "notProperty"}, env)
		},
	)

	accesssNonObject := ast.Property{ast.BooleanLiteral{true}, "property"}
	assert.PanicsWithValue(t,
		typeError("object", PawnBoolean{true}),
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
	assert.PanicsWithValue(t, variableNotFoundError("notVar"), func() {
		evaluateMemberName(memberNameNotFound, env)
	})
	assert.PanicsWithValue(t, variableNotFoundError("notVar"), func() {
		evaluateMemberName(memberNameNotFound, envDeep)
	})

	// Var exists, shallow and deep search:
	assert.Equal(t, PropertyOrVar{env.variables, "var"}, evaluateMemberName(memberNameFound, env))
	assert.Equal(t, PropertyOrVar{env.variables, "var"}, evaluateMemberName(memberNameFound, envDeep))
}
