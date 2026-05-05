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
		t, numberOperation(PawnInt{0}, PawnInt{0}, dummyIntHandler, dummyFloatHandler),	PawnInt{0},
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
	};

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

