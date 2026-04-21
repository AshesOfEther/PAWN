package main

// Value in the semantics is, as of writing, defined as
// V = Loc ∪ Object ∪ Int ∪ Float ∪ String ∪ List ∪ Function ∪ Boolean

// In go, this sort of "∪" type can be done with an interface
// where each of the types connected by ∪ is a struct that implements the interface.
// The interface is just a single unimportant function called value.
// To be clear, every sub-type (e.g StringP) has to implement the function "value()", and they "automatically" implement the interface.

// Source for this strategy: https://appliedgo.net/spotlight/sum-types-in-go/

// Also note, that instead of a generic Loc that's described in the semantics,
// and instead of objects being a value that isn't behind a reference,
// this implementation uses pointers directly to only lists or objects.
// The behavior should be equivalent. If it isn't, the semantics is likely wrong/unintended.

type valueP interface {
	value()
}

//--------
// Object
//--------
type objectP struct {
	v *struct{
		// TODO
	}
}
func (l objectP) value() {}

//--------
// Int
//--------
type intP struct {
	v int64
}
func (l intP) value() {}

//--------
// Float
//--------
type floatP struct {
	v float64
}
func (l floatP) value() {}

//--------
// String
//--------
type stringP struct {
	v string
}
func (l stringP) value() {}

//--------
// List
//--------
type listP struct {
	v *struct{
		// TODO
	}
}
func (l listP) value() {}

//--------
// Function
//--------
type functionP struct {
	// TODO
}
func (l functionP) value() {}

//--------
// Boolean
//--------
type booleanP struct {
	v bool
}
func (l booleanP) value() {}
