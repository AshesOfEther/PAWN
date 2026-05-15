package interpreter

import (
	"fmt"

	"pawn/ast"
)

type ListOrIndexOrPropertyOrVar interface {
	ListOrIndexOrPropertyOrVar()
}

type Index struct {
	// Not pointer to list.
	// List only re-alocates when changing size, and
	// the list cannot change size while MemberReturnedIndex is in use.
	list []PawnValue
	index int64
}

type PropertyOrVar struct {
	map_ map[string]PawnValue
	key string
}

func (PawnList) ListOrIndexOrPropertyOrVar() {}
func (Index) ListOrIndexOrPropertyOrVar() {}
func (PropertyOrVar) ListOrIndexOrPropertyOrVar() {}

func evaluateMember(member ast.Member, environment Environment) ListOrIndexOrPropertyOrVar {
	switch member := member.(type) {
		case ast.Name:
			return evaluateMemberName(member, environment)
		case ast.Property:
			return evaluateMemberProperty(member, environment)
		case ast.Index:
			return evaluateMemberIndex(member, environment)
		default:
			panic(fmt.Sprint("Unexpected invalid Member: ", member))
	}
}

func evaluateMemberName(member ast.Name, environment Environment) ListOrIndexOrPropertyOrVar {
	if _, ok := environment.variables[member.Name]; ok {
		return PropertyOrVar{environment.variables, member.Name}
	}
	if environment.parent != nil {
		return evaluateMemberName(member, *environment.parent)
	}
	panic(PawnError{
		fmt.Sprint("Variable '", member.Name, "' doesn't exist"),
	})
}

func evaluateMemberProperty(member ast.Property, environment Environment) ListOrIndexOrPropertyOrVar {
	object, ok := EvaluateExpression(member.Object, environment).(PawnObject)
	if !ok {
		panic(PawnError{
			fmt.Sprint("Tried accessing property of non-object: ", object),
		})
	}
	if _, ok := object.v[member.Property]; ok {
		return PropertyOrVar{object.v, member.Property}
	}
	panic(PawnError{
		fmt.Sprint("Field '", member.Property, "' doesn't exist in object"),
	})
}

func evaluateMemberIndex(member ast.Index, environment Environment) ListOrIndexOrPropertyOrVar {
	maybeList := EvaluateExpression(member.List, environment)
	pawnList, ok := maybeList.(PawnList)
	if !ok {
		panic(PawnError{
			fmt.Sprint("Tried indexing non-list: ", maybeList),
		})
	}

	if len(member.Index) == 1 {
		maybeIndex := EvaluateExpression(member.Index[0], environment)
		index, ok := maybeIndex.(PawnInt)

		if !ok {
			panic(PawnError{
				fmt.Sprint("Cannot index with a non-integer: ", maybeIndex),
			})
		}
		validateIndex(index.v, pawnList)
		return Index{*pawnList.v, index.v}
	}

	indecies := make([]int64, len(member.Index))
	for i, v := range member.Index {
		maybeIndex := EvaluateExpression(v, environment)

		indexPawn, ok := maybeIndex.(PawnInt)
		if !ok {
			panic(PawnError{
				fmt.Sprint("Tried multi-indexing but value at position ",i , " was a non-integer: ", maybeIndex),
			})
		}
		indecies[i] = indexPawn.v
	}

	outputList := make([]PawnValue, len(member.Index))
	for i, v := range indecies {
		validateIndex(v, pawnList)
		outputList[i] = (*pawnList.v)[v]
	}

	return PawnList{&outputList}
}

func validateIndex(index int64, pawnList PawnList) {
	if index < 0 {
		panic(PawnError{
			fmt.Sprint("Cannot index with a negative integer: ", index),
		})
	}
	if index >= int64(len(*pawnList.v)) {
		panic(PawnError{
			fmt.Sprint("Cannot index with ", index, " because it is not less than ", len(*pawnList.v),", the list's size"),
		})
	}
}
