package interpreter

import "pawn/ast"

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
type BasicValue[T any] struct {
	v T
}

type PawnBoolean BasicValue[bool]
type PawnObject BasicValue[map[string]*PawnValue] // The map and pointer shouldn't be nil
type PawnInt BasicValue[int64]
type PawnFloat BasicValue[float64]
type PawnString BasicValue[string]
type PawnList BasicValue[*[]*PawnValue] // Both pointers here shouldn't be nil

func (l PawnBoolean) value() {}
func (l PawnObject) value() {}
func (l PawnInt) value() {}
func (l PawnFloat) value() {}
func (l PawnString) value() {}
func (l PawnList) value() {}

//--------
// Function
//--------
type PawnFunctionPrimitive struct {
	f func([]PawnValue) *PawnValue // Possibly nil pointer, incase there's no returned value
	positionalArgCount uint
	namedArguments map[string]ast.Expression // This map shouldn't be nil
}
func (l PawnFunctionPrimitive) value() {}

type PawnFunctionUser struct {
	environment Environment
	positionalArguments []string
	namedArguments map[string]ast.Expression // This map shouldn't be nil
	body ast.Statement
}
func (l PawnFunctionUser) value() {}
