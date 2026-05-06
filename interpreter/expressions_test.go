package interpreter

import (
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

// Testing for all arithmetic operations
func TestAdd(t *testing.T) {
	assert.Equal(t, evaluateAdd(PawnInt{1}, PawnInt{1}), PawnInt{2})
	assert.Equal(t, evaluateAdd(PawnInt{1}, PawnFloat{1.1}), PawnFloat{2.1})
	assert.Equal(t, evaluateAdd(PawnFloat{1.1}, PawnFloat{1.1}), PawnFloat{2.2})
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, evaluateSubtract(PawnInt{1}, PawnInt{1}), PawnInt{0})
	assert.Equal(t, evaluateSubtract(PawnInt{1}, PawnFloat{2.1}), PawnFloat{-1.1})
	assert.Equal(t, evaluateSubtract(PawnFloat{1.1}, PawnFloat{1.1}), PawnFloat{0})
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, evaluateMultiply(PawnInt{1}, PawnInt{1}), PawnInt{1})
	assert.Equal(t, evaluateMultiply(PawnInt{1}, PawnFloat{2.1}).(PawnFloat).v, 2.1)
	assert.InEpsilon(t, 1.21, evaluateMultiply(PawnFloat{1.1}, PawnFloat{1.1}).(PawnFloat).v, 0.001)
}

func TestDivide(t *testing.T) {
	assert.Equal(t, evaluateDivide(PawnInt{10}, PawnInt{5}), PawnInt{2})
	assert.InEpsilon(t, 0.4761904761904762, evaluateDivide(PawnInt{1}, PawnFloat{2.1}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 2.058823529411765, evaluateDivide(PawnFloat{10.5}, PawnFloat{5.1}).(PawnFloat).v, 0.001)
}

func TestModulo(t *testing.T) {
	assert.Equal(t, evaluateModulo(PawnInt{1}, PawnInt{1}), PawnInt{0})
	assert.InEpsilon(t, 10, evaluateModulo(PawnInt{10}, PawnFloat{100}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 10, evaluateModulo(PawnFloat{10}, PawnFloat{100}).(PawnFloat).v, 0.001)
}

func TestPower(t *testing.T) {
	assert.Equal(t, evaluatePower(PawnInt{10}, PawnInt{5}), PawnInt{100000})
	assert.Equal(t, evaluatePower(PawnInt{0}, PawnInt{0}), PawnInt{1})
	assert.InEpsilon(t, 1, evaluatePower(PawnInt{1}, PawnFloat{2.1}).(PawnFloat).v, 0.001)
	assert.InEpsilon(t, 161460.17737759, evaluatePower(PawnFloat{10.5}, PawnFloat{5.1}).(PawnFloat).v, 0.001)
}
