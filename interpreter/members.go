package interpreter

import (
	"fmt"
	"slices"

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
	list []PawnValue
}

type MemberReturnedIndex struct {
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

				if indexPawn, ok := maybeIndex.(PawnInt); ok {
					return MemberReturnedIndex{slices.Clone(*listPawn.v), indexPawn.v}
				} else {
					panic(PawnError{
						fmt.Sprint("Tried indexing with non-integer: ", maybeIndex),
					})
				}
			} else {
				indices := make([]int64, len(member.Index))

				for i, v := range member.Index {
					maybeIndex := EvaluateExpression(v, environment)

					if indexPawn, ok := maybeIndex.(PawnInt); ok {
						indices[i] = indexPawn.v
					} else {
						panic(PawnError{
							fmt.Sprint("Tried indexing with non-integer: ", maybeIndex),
						})
					}
				}

				outputList := make([]PawnValue, len(indices))

				for i, v := range indices {
					outputList[i] = (*listPawn.v)[v]
				}

				return MemberReturnedList{outputList}
			}

		default:
			panic(fmt.Sprint("Unexpected invalid Member: ", member))
	}
}
