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
