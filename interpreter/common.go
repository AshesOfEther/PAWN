package interpreter

import (
	"fmt"
	"strings"
)

type Environment struct {
	variables map[string]PawnValue
	parent *Environment // Can be nil
}

func NewEnvironment(parent *Environment) Environment {
	return Environment{
		map[string]PawnValue{},
		parent,
	}
}

func pawnPrint(s PawnString) {
	println(s.v)
}

func CreateGlobalScope() Environment {
	environment := NewEnvironment(nil)
	environment.variables["print"] = newPrimitiveFunction(pawnPrint)
	return environment
}

type PawnError struct {
	Message string
}

func typeError(expectedType string, gotValue PawnValue) PawnError {
	gotType := "unknown"
	switch gotValue.(type) {
		case PawnObject:
			gotType = "object"
		case PawnBoolean:
			gotType = "boolean"
		case PawnInt:
			gotType = "integer"
		case PawnFloat:
			gotType = "float"
		case PawnString:
			gotType = "string"
		case PawnList:
			gotType = "list"
		case PawnFunctionPrimitive, PawnFunctionUser:
			gotType = "function"
	}
	return PawnError{
		fmt.Sprintf("expected %s, got %s: %v", expectedType, gotType, gotValue),
	}
}

func duplicateNamedArgumentDeclarationError(name string) PawnError {
	return PawnError{
		fmt.Sprintf("named argument '%s' is declared more than once", name),
	}
}

func mismatchedArgumentCountError(expectedCount int, gotCount int) PawnError {
	return PawnError{
		fmt.Sprintf(
			"function expects %d positional arguments, but %d were provided",
			expectedCount, gotCount,
		),
	}
}

func unexpectedNamedArgumentsError(unexpectedNames []string) PawnError {
	argumentString := "arguments were"
	if len(unexpectedNames) == 1 {
		argumentString = "argument was"
	}
	return PawnError{
		fmt.Sprintf(
			"unexpected named %s provided to function: %s",
			argumentString, strings.Join(unexpectedNames, ", "),
		),
	}
}

func usedVoidReturnValue() PawnError {
	return PawnError{
		"tried to use return value of function call that did not return anything",
	}
}

func negativePowerError(gotPower int64) PawnError {
	return PawnError{
		fmt.Sprintf("expected non-negative power for integer exponentiation, got %v", gotPower),
	}
}

func variableNotFoundError(name string) PawnError {
	return PawnError{
		fmt.Sprintf("tried to access undefined variable '%s'", name),
	}
}

func negativeIndexError(index int64) PawnError {
	return PawnError{
		fmt.Sprintf("Cannot index with a negative integer: %v", index),
	}
}

func indexTooHighError(index int64, pawnList PawnList) PawnError {
	return PawnError{
		fmt.Sprintf("Cannot index with %v because it is not less than %v, the list's size", index, len(*pawnList.v)),
	}
}

func fieldNonExistantError(property string) PawnError {
	return PawnError{
		fmt.Sprintf("Property '%s' doesn't exist in object", property),
	}
}

func divideByZeroError() PawnError {
	return PawnError{
		("expected non-zero divisor for integer division, got zero"),
	}
}

func rangeFloatError() PawnError {
	return PawnError{
		("Range was given float"),
	}
}

func unexpectedBodyError() PawnError {
	return PawnError{
		"function does not expect a body, but it was called with one",
	}
}
