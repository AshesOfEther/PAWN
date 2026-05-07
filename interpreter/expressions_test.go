package interpreter

import (
	"pawn/ast"
	"testing"
	"math"

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

func TestEvaluateEqual(t *testing.T) {

	// Except for floats, the total amount of cases is 3*amountOfTypes
	// Two values T1 T2 of each type T to compare 3 cases,
	// where T1 != T2 in the most non-trivial way.
	// 1. a1 == a1 is true
	// 2. a1 == a2 is false
	// 3. a1 == b1 is false

	emptyStringList := []string{}
	PrimitiveFunctionForTest := func (x []PawnValue) *PawnValue {
		return nil
	}
	env := NewEnvironment(nil)
	emptyBody := []ast.Statement{}
	noOptionalArgs := map[string]ast.Expression{}

	// Values
	B1 := PawnBoolean{true}
	B2 := PawnBoolean{false} // false since that's the only other case of this type
	O1 := PawnObject{map[string]PawnValue{}}
	O2 := PawnObject{map[string]PawnValue{}} // Maps are equal in content but have different reference
	F1 := PawnFloat{0.3}
	F2 := PawnFloat{float64(0.1)+float64(0.2)} // Floats round
	L1 := PawnList{&[]PawnValue{}}
	L2 := PawnList{&[]PawnValue{}} // Lists are equal in content but have different reference
	P1 := PawnFunctionPrimitive{&PawnFunctionPrimitiveInner{f: PrimitiveFunctionForTest, positionalArgCount: 0, namedArguments: emptyStringList}}
	P2 := PawnFunctionPrimitive{&PawnFunctionPrimitiveInner{f: PrimitiveFunctionForTest, positionalArgCount: 0, namedArguments: emptyStringList}} // Functions are equal in content but have different reference
	U1 := PawnFunctionUser{&PawnFunctionUserInner{environment: env, positionalArguments: emptyStringList, namedArguments: noOptionalArgs, body: emptyBody}}
	U2 := PawnFunctionUser{&PawnFunctionUserInner{environment: env, positionalArguments: emptyStringList, namedArguments: noOptionalArgs, body: emptyBody}}

	// Float nans are special, they aren't equal to themselves
	F3 := PawnFloat{math.NaN()}

	// Bool
	assert.True(t,  evaluateEqual(B1, B1).(PawnBoolean).v)  // a1 == a1
	assert.True(t, !evaluateEqual(B1, B2).(PawnBoolean).v)  // a1 != a2
	assert.True(t, !evaluateEqual(B1, U1).(PawnBoolean).v)  // a1 != b1

	// Object
	assert.True(t,  evaluateEqual(O1, O1).(PawnBoolean).v)  // a1 == a1
	assert.True(t, !evaluateEqual(O1, O2).(PawnBoolean).v)  // a1 != a2
	assert.True(t, !evaluateEqual(O1, B1).(PawnBoolean).v)  // a1 != b1

	// Float
	assert.True(t,  evaluateEqual(F1, F1).(PawnBoolean).v)  // a1 == a1
	assert.True(t, !evaluateEqual(F1, F2).(PawnBoolean).v)  // a1 != a2
	assert.True(t, !evaluateEqual(F1, O1).(PawnBoolean).v)  // a1 != b1
	assert.True(t, !evaluateEqual(F3, F3).(PawnBoolean).v)  // a1 == a1

	// List
	assert.True(t,  evaluateEqual(L1, L1).(PawnBoolean).v)  // a1 == a1
	assert.True(t, !evaluateEqual(L1, L2).(PawnBoolean).v)  // a1 != a2
	assert.True(t, !evaluateEqual(L1, F1).(PawnBoolean).v)  // a1 != b1

	// Primtitive
	assert.True(t,  evaluateEqual(P1, P1).(PawnBoolean).v)  // a1 == a1
	assert.True(t, !evaluateEqual(P1, P2).(PawnBoolean).v)  // a1 != a2
	assert.True(t, !evaluateEqual(P1, L1).(PawnBoolean).v)  // a1 != b1

	// User
	assert.True(t,  evaluateEqual(U1, U1).(PawnBoolean).v)  // a1 == a1
	assert.True(t, !evaluateEqual(U1, U2).(PawnBoolean).v)  // a1 != a2
	assert.True(t, !evaluateEqual(U1, P1).(PawnBoolean).v)  // a1 != b1
}
