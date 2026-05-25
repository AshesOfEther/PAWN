package interpreter

import (
	"reflect"
	
	"pawn/ast"
)

// ----------------------------------
// # Value definition
// ----------------------------------

// Value in the semantics is, as of writing, defined as
// V = Loc ∪ Object ∪ Int ∪ Float ∪ String ∪ List ∪ Function ∪ Boolean

// In go, this sort of "∪" type can be done with an interface
// where each of the types connected by ∪ is a struct that implements the interface.
// The interface is just a single unimportant function called value.
// To be clear, every sub-type (e.g PawnString) has to implement the function "value()", and they "automatically" implement the interface.

// Source for this strategy: https://appliedgo.net/spotlight/sum-types-in-go/

// Also note, that instead of a generic Loc that's described in the semantics,
// and instead of objects being a value that isn't behind a reference,
// this implementation uses pointers directly to only lists or objects.
// The behavior should be equivalent. If it isn't, the semantics is likely wrong/unintended.

type PawnValue interface {
	value()
}

// Used to define Bool, Int, Float, String, List, and Object.
type BasicPawnValue[T any] struct {
	v T
}

type PawnBoolean BasicPawnValue[bool]
type PawnObject BasicPawnValue[map[string]PawnValue]
type PawnInt BasicPawnValue[int64]
type PawnFloat BasicPawnValue[float64]
type PawnString BasicPawnValue[string]
type PawnList BasicPawnValue[*[]PawnValue]

func (l PawnBoolean) value() {}
func (l PawnObject) value() {}
func (l PawnInt) value() {}
func (l PawnFloat) value() {}
func (l PawnString) value() {}
func (l PawnList) value() {}

//--------
// Function
//--------
type PawnFunctionPrimitiveInner struct {
	// Stores a function with the following arguments:
	//
	// - Zero or more positional arguments: Either `PawnValue` or a type that implements it
	// - (Optional) A struct representing the named arguments, with fields of type `PawnValue` or `*T`
	//   where `T` implements it
	// - (Optional) A `map[string]PawnValue` containing variables defined in the body of the call
	//
	// If present, the return type must either be `PawnValue`, `T`, or `*T`, where `T` implements
	// `PawnValue`.
	// 
	// Positional arguments must not be nil.
	f reflect.Value
	positionalArgTypes []reflect.Type
	namedArgumentStructType reflect.Type // Is nil if the underlying function doesn't accept a named arguments struct
	namedArgumentTypes map[string]reflect.Type // Is nil if the underlying function doesn't accept a named arguments struct
	acceptsBody bool
}
type PawnFunctionUserInner struct {
	environment Environment
	positionalArguments []string
	namedArguments map[string]ast.Expression
	body []ast.Statement
}

type PawnFunctionPrimitive BasicPawnValue[*PawnFunctionPrimitiveInner]
type PawnFunctionUser BasicPawnValue[*PawnFunctionUserInner]

func (l PawnFunctionPrimitive) value() {}
func (l PawnFunctionUser) value() {}
