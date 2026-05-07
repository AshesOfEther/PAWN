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

func negativePowerError(gotPower int64) PawnError {
	return PawnError{
		fmt.Sprintf("expected non-negative power for integer exponentiation, got %v", gotPower),
	}
}

func divideByZeroError() PawnError {
	return PawnError{
		("expected non-zero divisor for integer division, got zero"),
	}
}
