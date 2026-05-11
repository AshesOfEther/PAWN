package interpreter

import (
	"pawn/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dummyIntHandler(_ int64, _ int64) PawnValue {
	return PawnInt{0}
}
func dummyFloatHandler(_ float64, _ float64) PawnValue {
	return PawnFloat{0}
}

func TestGreaterThanInt(t *testing.T) {
	assert.Equal(t, evaluateGreaterThan(PawnInt{2}, PawnInt{1}), PawnBoolean{true})
	assert.Equal(t, evaluateGreaterThan(PawnInt{1}, PawnInt{2}), PawnBoolean{false})
	assert.Equal(t, evaluateGreaterThan(PawnInt{1}, PawnInt{1}), PawnBoolean{false})
}

func TestGreaterThanFloat(t *testing.T) {
	assert.Equal(t, evaluateGreaterThan(PawnFloat{2}, PawnFloat{1}), PawnBoolean{true})
	assert.Equal(t, evaluateGreaterThan(PawnFloat{1}, PawnFloat{2}), PawnBoolean{false})
	assert.Equal(t, evaluateGreaterThan(PawnFloat{1}, PawnFloat{1}), PawnBoolean{false})
}

func TestLessThanInt(t *testing.T)  {
	assert.Equal(t, evaluateLessThan(PawnInt{2}, PawnInt{1}), PawnBoolean{false})
	assert.Equal(t, evaluateLessThan(PawnInt{1}, PawnInt{2}), PawnBoolean{true})
	assert.Equal(t, evaluateLessThan(PawnInt{1}, PawnInt{1}), PawnBoolean{false})
}

func TestLessThanFloat(t *testing.T)  {
	assert.Equal(t, evaluateLessThan(PawnFloat{2}, PawnFloat{1}), PawnBoolean{false})
	assert.Equal(t, evaluateLessThan(PawnFloat{1}, PawnFloat{2}), PawnBoolean{true})
	assert.Equal(t, evaluateLessThan(PawnFloat{1}, PawnFloat{1}), PawnBoolean{false})
}

func TestLessOrEqualInt(t *testing.T)  {
	assert.Equal(t, evaluateLessOrEqual(PawnInt{2}, PawnInt{1}), PawnBoolean{false})
	assert.Equal(t, evaluateLessOrEqual(PawnInt{1}, PawnInt{2}), PawnBoolean{true})
	assert.Equal(t, evaluateLessOrEqual(PawnInt{1}, PawnInt{1}), PawnBoolean{true})
}

func TestLessOrEqualFloat(t *testing.T)  {
	assert.Equal(t, evaluateLessOrEqual(PawnFloat{2}, PawnFloat{1}), PawnBoolean{false})
	assert.Equal(t, evaluateLessOrEqual(PawnFloat{1}, PawnFloat{2}), PawnBoolean{true})
	assert.Equal(t, evaluateLessOrEqual(PawnFloat{1}, PawnFloat{1}), PawnBoolean{true})
}

func TestGreaterOrEqualInt(t *testing.T)  {
	assert.Equal(t, evaluateGreaterOrEqual(PawnInt{2}, PawnInt{1}), PawnBoolean{true})
	assert.Equal(t, evaluateGreaterOrEqual(PawnInt{1}, PawnInt{2}), PawnBoolean{false})
	assert.Equal(t, evaluateGreaterOrEqual(PawnInt{1}, PawnInt{1}), PawnBoolean{true})
}

func TestGreaterOrEqualFloat(t *testing.T)  {
	assert.Equal(t, evaluateGreaterOrEqual(PawnFloat{2}, PawnFloat{1}), PawnBoolean{true})
	assert.Equal(t, evaluateGreaterOrEqual(PawnFloat{1}, PawnFloat{2}), PawnBoolean{false})
	assert.Equal(t, evaluateGreaterOrEqual(PawnFloat{1}, PawnFloat{1}), PawnBoolean{true})
}

func TestNumberOperationCoerceIntToFloat(t *testing.T) {
	assert.Equal(
		t, numberOperation(PawnInt{0}, PawnInt{0}, dummyIntHandler, dummyFloatHandler), PawnInt{0},
	)
	assert.Equal(
		t, numberOperation(PawnFloat{0}, PawnInt{0}, dummyIntHandler, dummyFloatHandler), PawnFloat{0},
	)
	assert.Equal(
		t, numberOperation(PawnInt{0}, PawnFloat{0}, dummyIntHandler, dummyFloatHandler), PawnFloat{0},
	)
	assert.Equal(
		t, numberOperation(PawnFloat{0}, PawnFloat{0}, dummyIntHandler, dummyFloatHandler), PawnFloat{0},
	)
}

func TestNumberOperationRejectNonNumber(t *testing.T) {
	// Only the type is important here. The fields don't need to be valid.
	values := []PawnValue{
		PawnBoolean{},
		PawnObject{},
		PawnString{},
		PawnList{},
		PawnFunctionPrimitive{},
		PawnFunctionUser{},
	}

	for _, value := range values {
		panicValue := typeError("int or float", value)

		assert.PanicsWithValue(t, panicValue, func() {
			numberOperation(PawnInt{0}, value, dummyIntHandler, dummyFloatHandler)
		})
		assert.PanicsWithValue(t, panicValue, func() {
			numberOperation(PawnFloat{0}, value, dummyIntHandler, dummyFloatHandler)
		})
		assert.PanicsWithValue(t, panicValue, func() {
			numberOperation(value, PawnInt{0}, dummyIntHandler, dummyFloatHandler)
		})
		assert.PanicsWithValue(t, panicValue, func() {
			numberOperation(value, PawnFloat{0}, dummyIntHandler, dummyFloatHandler)
		})
		assert.PanicsWithValue(t, panicValue, func() {
			numberOperation(value, value, dummyIntHandler, dummyFloatHandler)
		})
	}
}

func TestEvaluateAnd(t *testing.T) {
	assert.Equal(t, evaluateAnd(PawnBoolean{true}, PawnBoolean{true}), PawnBoolean{true})
	assert.Equal(t, evaluateAnd(PawnBoolean{true}, PawnBoolean{false}), PawnBoolean{false})
	assert.Equal(t, evaluateAnd(PawnBoolean{false}, PawnBoolean{true}), PawnBoolean{false})
	assert.Equal(t, evaluateAnd(PawnBoolean{false}, PawnBoolean{false}), PawnBoolean{false})
}

func TestEvaluateOr(t *testing.T) {
	assert.Equal(t, evaluateOr(PawnBoolean{true}, PawnBoolean{true}), PawnBoolean{true})
	assert.Equal(t, evaluateOr(PawnBoolean{true}, PawnBoolean{false}), PawnBoolean{true})
	assert.Equal(t, evaluateOr(PawnBoolean{false}, PawnBoolean{true}), PawnBoolean{true})
	assert.Equal(t, evaluateOr(PawnBoolean{false}, PawnBoolean{false}), PawnBoolean{false})
}

func TestErrorOnCallingNonFunction(t *testing.T) {
	assert.PanicsWithValue(t, typeError("function", PawnInt{0}), func() {
		evaluateFunctionCall(ast.FunctionCall{
			ast.IntLiteral{0},
			[]ast.Expression{},
			[]ast.NamedArg{},
			[]ast.Statement{},
		}, NewEnvironment(nil))
	})
}

func dummyValidateNamedArg(name string) bool {
	return name == "foo" || name == "bar"
}

var oneNamedArg []ast.NamedArg = []ast.NamedArg{
	ast.NamedArg{"foo", ast.IntLiteral{0}},
}

func TestValidateArgumentsCorrect(t *testing.T) {
	validateFunctionArguments(3, 3, oneNamedArg, dummyValidateNamedArg)
}

func TestValidateArgumentsNotEnoughPositional(t *testing.T) {
	assert.PanicsWithValue(t, mismatchedArgumentCountError(3, 1), func() {
		validateFunctionArguments(1, 3, oneNamedArg, dummyValidateNamedArg)
	})
}

func TestValidateArgumentsTooManyPositional(t *testing.T) {
	assert.PanicsWithValue(t, mismatchedArgumentCountError(3, 5), func() {
		validateFunctionArguments(5, 3, oneNamedArg, dummyValidateNamedArg)
	})
}

func TestValidateArgumentsUnexpectedNamedArg(t *testing.T) {
	assert.PanicsWithValue(t, unexpectedNamedArgumentsError([]string{"baz"}), func() {
		validateFunctionArguments(3, 3, []ast.NamedArg{
			ast.NamedArg{"foo", ast.IntLiteral{0}},
			ast.NamedArg{"bar", ast.IntLiteral{1}},
			ast.NamedArg{"baz", ast.IntLiteral{2}},
		}, dummyValidateNamedArg)
	})
}

func TestMemberIndexInExpression(t *testing.T) {
	env := NewEnvironment(nil)
	env.variables["varForList"] = PawnList{
		&[]PawnValue{PawnBoolean{true},
		PawnList{&[]PawnValue{}}},
	}

	indexInExpression := ast.MemberExpression{ast.Index{
		ast.MemberExpression{
			ast.Name{"varForList"},
		},
		[]ast.Expression{ast.IntLiteral{1}},
	}}
	assert.Equal(t,
		PawnList{&[]PawnValue{}},
		evaluateMemberExpression(indexInExpression, env),
	)
}

func TestMemberIndeciesInExpression(t *testing.T) {
	env := NewEnvironment(nil)
	env.variables["varForList"] = PawnList{
		&[]PawnValue{
			PawnBoolean{true},
			PawnBoolean{false},
		},
	}

	indexInExpression := ast.MemberExpression{ast.Index{
		ast.MemberExpression{
			ast.Name{"varForList"},
		},
		[]ast.Expression{ast.IntLiteral{1}, ast.IntLiteral{0}},
	}}
	assert.Equal(t,
		PawnList{&[]PawnValue{PawnBoolean{false}, PawnBoolean{true}}},
		evaluateMemberExpression(indexInExpression, env),
	)
}

func TestMemberPropertyInExpression(t *testing.T) {
	env := NewEnvironment(nil)
	env.variables["varForObj"] = PawnObject{
		map[string]PawnValue{
			"one": PawnBoolean{true},
			"two": PawnBoolean{false},
		},
	}

	indexInExpression := ast.MemberExpression{ast.Property{
		ast.MemberExpression{
			ast.Name{"varForObj"},
		},
		"one",
	}}
	assert.Equal(t,
		PawnBoolean{true},
		evaluateMemberExpression(indexInExpression, env),
	)
}


// Testing for all arithmetic operations
func TestAdd(t *testing.T) {
	assert.Equal(t, evaluateAdd(PawnInt{1}, PawnInt{1}), PawnInt{2})
	assert.InEpsilon(t, 2.1, evaluateAdd(PawnInt{1}, PawnFloat{1.1}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 2.2, evaluateAdd(PawnFloat{1.1}, PawnFloat{1.1}).(PawnFloat).v, 0.001)
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, evaluateSubtract(PawnInt{1}, PawnInt{1}), PawnInt{0})
	assert.InEpsilon(t, -1.1, evaluateSubtract(PawnInt{1}, PawnFloat{2.1}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 1, evaluateSubtract(PawnFloat{2.1}, PawnFloat{1.1}).(PawnFloat).v, 0.001)
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, evaluateMultiply(PawnInt{1}, PawnInt{1}), PawnInt{1})
	assert.InEpsilon(t, 2.1, evaluateMultiply(PawnInt{1}, PawnFloat{2.1}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 1.21, evaluateMultiply(PawnFloat{1.1}, PawnFloat{1.1}).(PawnFloat).v, 0.001)
}

func TestDivide(t *testing.T) {
	assert.Equal(t, evaluateDivide(PawnInt{10}, PawnInt{5}), PawnInt{2})
	assert.InEpsilon(t, 0.4761904761904762, evaluateDivide(PawnInt{1}, PawnFloat{2.1}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 2.058823529411765, evaluateDivide(PawnFloat{10.5}, PawnFloat{5.1}).(PawnFloat).v, 0.001)
}

//For negative int and float deivisor go uses truncated devision
func TestModulo(t *testing.T) {
	assert.Equal(t, evaluateModulo(PawnInt{1}, PawnInt{1}), PawnInt{0})
	assert.Equal(t, evaluateModulo(PawnInt{-1}, PawnInt{2}), PawnInt{-1})
	assert.Equal(t, evaluateModulo(PawnInt{1}, PawnInt{-2}), PawnInt{1})
	assert.InEpsilon(t, 2, evaluateModulo(PawnInt{5}, PawnFloat{3}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 2, evaluateModulo(PawnFloat{5}, PawnFloat{3}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, -2, evaluateModulo(PawnFloat{-5}, PawnFloat{3}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 2, evaluateModulo(PawnFloat{5}, PawnFloat{-3}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, -2, evaluateModulo(PawnFloat{-5}, PawnFloat{-3}).(PawnFloat).v, 0.001)
}

func TestPower(t *testing.T) {
	assert.Equal(t, evaluatePower(PawnInt{10}, PawnInt{5}), PawnInt{100000})
	assert.Equal(t, evaluatePower(PawnInt{0}, PawnInt{0}), PawnInt{1})
	assert.InEpsilon(t, 1, evaluatePower(PawnInt{1}, PawnFloat{2.1}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 161460.17737759, evaluatePower(PawnFloat{10.5}, PawnFloat{5.1}).(PawnFloat).v, 0.001)
}

func TestArithmeticErrors(t *testing.T)  {
	assert.PanicsWithValue(t, divideByZeroError(), func() {evaluateDivide(PawnInt{1}, PawnInt{0})})
	assert.PanicsWithValue(t, negativePowerError(-1), func() {evaluatePower(PawnInt{1}, PawnInt{-1})})
}

//Testing for range operations
func TestRange(t *testing.T) {
	assert.Equal(t, evaluateRange(PawnInt{0}, PawnInt{3}), PawnList{&[]PawnValue{PawnInt{0},PawnInt{1},PawnInt{2}}})
	assert.Equal(t, evaluateRange(PawnInt{0}, PawnInt{0}), PawnList{&[]PawnValue{}})
}

func TestRangeInclusive(t *testing.T) {
	assert.Equal(t, evaluateRangeInclusive(PawnInt{0}, PawnInt{3}), PawnList{&[]PawnValue{PawnInt{0},PawnInt{1},PawnInt{2}, PawnInt{3}}})
	assert.Equal(t, evaluateRangeInclusive(PawnInt{0}, PawnInt{0}), PawnList{&[]PawnValue{PawnInt{0}}})
}

func TestRangeErrorsOnFloat(t *testing.T)  {
	assert.PanicsWithValue(t, rangeFloatError(), func() {evaluateRange(PawnFloat{2}, PawnFloat{2})})
}

func TestRangeInclusiveErrorsOnFloat(t *testing.T) {
	assert.PanicsWithValue(t, rangeFloatError(), func() {evaluateRangeInclusive(PawnFloat{2}, PawnFloat{2})})
}

func TestEvaluateListEmpty(t *testing.T) {
	list := ast.List{[]ast.Expression{}}
	expected := PawnList{&[]PawnValue{}}
	assert.Equal(t, expected, evaluateList(list, NewEnvironment(nil)))
}

func TestEvaluateListNonEmpty(t *testing.T) {
	list := ast.List{[]ast.Expression{
		ast.IntLiteral{2},
		ast.BooleanLiteral{true},
	}}
	expected := PawnList{&[]PawnValue{
		PawnInt{2},
		PawnBoolean{true},
	}}
	assert.Equal(t, expected, evaluateList(list, NewEnvironment(nil)))
}
