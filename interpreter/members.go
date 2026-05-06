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
type MemberReturnedList BasicValue[[]*PawnValue]

func (_ MemberReturnedMap) memberReturnValue() {}
func (_ MemberReturnedList) memberReturnValue() {}

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
//		case ast.Property:
//			object, ok := EvaluateExpression(member.Object, environment).(PawnObject)
//			if !ok {
//				panic(PawnError{
//					fmt.Sprint("Tried accessing non-existant property: ", member.Property),
//				})
//			}
//			if _, ok := object.v[member.Property]; ok {
//				return object.v, member.Property
//			}
//			panic(PawnError{
//				fmt.Sprint("Field \"", member.Property, "\" doesn't exist in object"),
//			})
//		case ast.Index:
//			maybeListEvaluated := evaluateMember(member.List, environment)
//
//			var maybeList PawnValue
//			switch maybeListEvaluated := maybeListEvaluated.(type) {
//				case MemberReturnedMap:
//					maybeList = maybeListEvaluated.mutateMe[list.nameToMutateAt]
//				case MemberReturnedList:
//					maybeList = list.v
//			}
//
//			listPawn, ok := maybeList.(PawnList)
//			if !ok {
//				panic(PawnError{
//					fmt.Sprint("Tried indexing non-list: ", maybeList),
//				})
//			}
//			list := listPawn.v
//
//			indices := make([]int64, len(member.Index))
//
//			for i, v := range member.Index {
//				maybeIndex := EvaluateExpression(v, environment)
//				indexPawn, ok := maybeIndex.(PawnInt)
//
//				if !ok {
//					panic(PawnError{
//						fmt.Sprint("Tried indexing with non-integer: ", maybeIndex),
//					})
//				}
//
//				indices[i] = indexPawn.v
//			}
//
//			placesThatCanMutate := make([]*PawnValue, len(indices))
//
//			for i, v := range indices {
//				placesThatCanMutate[i] = &list[v]
//			}
//
//			return 

		default:
			panic(fmt.Sprint("Unexpected invalid Member: ", member))
	}
}
