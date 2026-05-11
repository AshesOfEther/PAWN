package interpreter

import (
	"fmt"

	"pawn/ast"
)

type MemberReturnValue interface {
	memberReturnValue()
}

type MemberReturnedMap struct {
	mutateMe map[string]PawnValue
	nameToMutateAt string
}

type MemberReturnedList struct {
	list PawnList
}

type MemberReturnedIndex struct {
	// Not pointer to list.
	// List only re-alocates when changing size, and
	// the list cannot change size while MemberReturnedIndex is in use.
	list []PawnValue
	indexToMutateAt int64
}

func (_ MemberReturnedMap) memberReturnValue() {}
func (_ MemberReturnedList) memberReturnValue() {}
func (_ MemberReturnedIndex) memberReturnValue() {}

func EvaluateMember(member ast.Member, environment Environment) MemberReturnValue {
	switch member := member.(type) {
		case ast.Name:
			if _, ok := environment.variables[member.Name]; ok {
				return MemberReturnedMap{environment.variables, member.Name}
			}
			if environment.parent != nil {
				return EvaluateMember(member, *environment.parent)
			}
			panic(PawnError{
				fmt.Sprint("Variable \"", member.Name, "\" doesn't exist"),
			})
		case ast.Property:
			object, ok := EvaluateExpression(member.Object, environment).(PawnObject)
			if !ok {
				panic(PawnError{
					fmt.Sprint("Tried accessing non-existant property: ", member.Property),
				})
			}
			if _, ok := object.v[member.Property]; ok {
				return MemberReturnedMap{object.v, member.Property}
			}
			panic(PawnError{
				fmt.Sprint("Field \"", member.Property, "\" doesn't exist in object"),
			})
		case ast.Index:
			maybeList := EvaluateExpression(member.List, environment)
			listPawn, ok := maybeList.(PawnList)
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
				if index.v < 0 {
					panic(PawnError{
						fmt.Sprint("Cannot index with a negative integer: ", index.v),
					})
				}
				if index.v >= int64(len(*listPawn.v)) {
					panic(PawnError{
						fmt.Sprint("Cannot index with the integer ", index.v, " because it is not less than the list's size ", len(*listPawn.v)),
					})
				}

				return MemberReturnedIndex{*listPawn.v, index.v}
			}

			indecies := make([]int64, len(member.Index))
			for i, v := range member.Index {
				maybeIndex := EvaluateExpression(v, environment)

				indexPawn, ok := maybeIndex.(PawnInt)
				if !ok {
					panic(PawnError{
						fmt.Sprint("Tried indexing with non-integer: ", maybeIndex),
					})
				}
				indecies[i] = indexPawn.v
			}

			outputList := make([]PawnValue, len(member.Index))
			for i, v := range indecies {
				if v < 0 {
					panic(PawnError{
						fmt.Sprint("Cannot index with a negative integer: ", v),
					})
				}
				if v >= int64(len(*listPawn.v)) {
					panic(PawnError{
						fmt.Sprint("Cannot index with the integer ", v, " because it is not less than the list's size ", len(*listPawn.v)),
					})
				}
				outputList[i] = (*listPawn.v)[v]
			}

			return MemberReturnedList{PawnList{&outputList}}
		default:
			panic(fmt.Sprint("Unexpected invalid Member: ", member))
	}
}
