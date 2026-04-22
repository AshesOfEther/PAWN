package main

// ----------------------------------
// # Syntax representations for values
//
// The abstract syntax tree could choose to represent expressions and statements
// in a way that's very unoptimal to use as values directly. Instead, these could be converted
// to the types under this header.
// ----------------------------------

type PawnEnvironmentAsValue map[string]*PawnValue
type PawnExpressionAsValue any // TODO
type PawnStatementAsValue any // TODO

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

//--------
// Boolean
//--------
type PawnBoolean struct {
	v bool
}
func (l PawnBoolean) value() {}

//--------
// Object
//--------
type PawnObject struct {
	v map[string]*PawnValue // The map and pointer here shouldn't be nil
}
func (l PawnObject) value() {}

//--------
// Int
//--------
type PawnInt struct {
	v int64
}
func (l PawnInt) value() {}

//--------
// Float
//--------
type PawnFloat struct {
	v float64
}
func (l PawnFloat) value() {}

//--------
// String
//--------
type PawnString struct {
	v string
}
func (l PawnString) value() {}

//--------
// List
//--------
type PawnList struct {
	v *[]*PawnValue // Both pointers here shouldn't be nil
}
func (l PawnList) value() {}

//--------
// Function
//--------
type PawnFunctionPrimitive struct {
	f func([]PawnValue) *PawnValue // Possibly nil pointer, incase there's no returned value
	positionalArguments []string
	optionalArguments map[string]PawnExpressionAsValue // This map shouldn't be nil
}
func (l PawnFunctionPrimitive) value() {}

type PawnFunctionUser struct {
	environment PawnEnvironmentAsValue
	positionalArguments []string
	optionalArguments map[string]PawnExpressionAsValue // This map shouldn't be nil
	body PawnStatementAsValue
}
func (l PawnFunctionUser) value() {}
