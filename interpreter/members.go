package interpreter

import (
	"fmt"

	"pawn/ast"
)

// Currently a resolved member can be one of:
//   List, Index, Property, or Var
type ResolvedMember interface {
	ResolvedMember()
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

func (PawnList) ResolvedMember() {}
func (Index) ResolvedMember() {}
func (PropertyOrVar) ResolvedMember() {}

func evaluateMember(member ast.Member, environment Environment) ResolvedMember {
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

func evaluateMemberName(member ast.Name, environment Environment) ResolvedMember {
	if _, ok := environment.variables[member.Name]; ok {
		return PropertyOrVar{environment.variables, member.Name}
	}
	if environment.parent != nil {
		return evaluateMemberName(member, *environment.parent)
	}
	panic(variableNotFoundError(member.Name))
}

func evaluateMemberProperty(member ast.Property, environment Environment) ResolvedMember {
	maybeObject := EvaluateExpression(member.Object, environment)
	object, ok := maybeObject.(PawnObject)
	if !ok {
		panic(typeError("object", maybeObject))
	}
	if _, ok := object.v[member.Property]; ok {
		return PropertyOrVar{object.v, member.Property}
	}
	panic(fieldNonExistantError(member.Property))
}

func evaluateMemberIndex(member ast.Index, environment Environment) ResolvedMember {
	maybeList := EvaluateExpression(member.List, environment)
	pawnList, ok := maybeList.(PawnList)
	if !ok {
		panic(typeError("list", maybeList))
	}

	if len(member.Index) == 0 {
		panic("Multi-indexing with 0 indecies is invalid")
	}

	if len(member.Index) == 1 {
		maybeIndex := EvaluateExpression(member.Index[0], environment)
		index, ok := maybeIndex.(PawnInt)

		if !ok {
			panic(typeError("int", maybeIndex))
		}
		if index.v < 0 || index.v >= int64(len(*pawnList.v)) {
			panic(indexError(index.v, pawnList))
		}
		return Index{*pawnList.v, index.v}
	}

	indecies := make([]int64, len(member.Index))
	for i, v := range member.Index {
		maybeIndex := EvaluateExpression(v, environment)

		indexPawn, ok := maybeIndex.(PawnInt)
		if !ok {
			panic(typeError("integer", maybeIndex))
		}
		indecies[i] = indexPawn.v
	}

	outputList := make([]PawnValue, len(member.Index))
	for i, index := range indecies {
		if index < 0 || index >= int64(len(*pawnList.v)) {
			panic(indexError(index, pawnList))
		}
		outputList[i] = (*pawnList.v)[index]
	}

	return PawnList{&outputList}
}
