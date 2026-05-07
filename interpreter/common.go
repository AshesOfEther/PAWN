package interpreter

import "fmt"

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
