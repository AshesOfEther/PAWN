package interpreter

import (
	"fmt"
	"slices"

	"pawn/ast"
)

// Errors when evaluating a function call without a return value. If this is
// undesired, the caller should call `evaluateFunctionCall` directly for function
// calls.
func EvaluateExpression(expression ast.Expression, environment Environment) PawnValue {
	switch expression := expression.(type) {
		case ast.BinaryOp:
			return evaluateBinaryOp(expression, environment)
		case ast.FunctionCall:
			result := evaluateFunctionCall(expression, environment)
			if result == nil {
				panic(usedVoidReturnValue())
			}
			return result
		case ast.List:
			panic("TODO")
		case ast.FloatLiteral:
			return PawnFloat{expression.Value}
		case ast.IntLiteral:
			return PawnInt{expression.Value}
		case ast.BooleanLiteral:
			return PawnBoolean{expression.Value}
		case ast.StringLiteral:
			return PawnString{expression.Value}
		case ast.MemberExpression:
			panic("TODO")
		default:
			panic(fmt.Sprintf("Unexpected invalid Expression: %t", expression))
	}
}

func evaluateBinaryOp(expression ast.BinaryOp, environment Environment) PawnValue {
	leftValue := EvaluateExpression(expression.Left, environment)
	rightValue := EvaluateExpression(expression.Right, environment)

	switch expression.Operator{
		case ast.Equal:
			return evaluateEqual(leftValue, rightValue)
		case ast.NotEqual:
			return evaluateNotEqual(leftValue, rightValue)
		case ast.GreaterThan:
			return evaluateGreaterThan(leftValue, rightValue)
		case ast.Add:
			panic("TODO")
		case ast.And:
			return evaluateAnd(leftValue, rightValue) 
		case ast.Divide:
			panic("TODO")
		case ast.GreaterOrEqual:
			panic("TODO")
		case ast.LessOrEqual:
			panic("TODO")
		case ast.LessThan:
			panic("TODO")
		case ast.Modulo:
			panic("TODO")
		case ast.Multiply:
			panic("TODO")
		case ast.Or:
			return evaluateOr(leftValue, rightValue)
		case ast.Power:
			panic("TODO")
		case ast.Range:
			panic("TODO")
		case ast.RangeInclusive:
			panic("TODO")
		case ast.Subtract:
			panic("TODO")
		default:
			panic(fmt.Sprintf("unexpected ast.Operator: %#v", expression.Operator))
	}
}


func evaluateEqual(leftValue PawnValue, rightValue PawnValue) PawnValue {
	result := false

	switch left := leftValue.(type) {
	case PawnBoolean:
		if right, ok := rightValue.(PawnBoolean); ok {
			result = left.v == right.v
		}
	case PawnObject:
		if right, ok := rightValue.(PawnObject); ok {
			result = &left == &right
		}
	case PawnInt:
		if right, ok := rightValue.(PawnInt); ok {
			result = left.v == right.v
		}
	case PawnFloat:
		if right, ok := rightValue.(PawnFloat); ok {
			result = left.v == right.v
		}
	case PawnString:
		if right, ok := rightValue.(PawnString); ok {
			result = left.v == right.v
		}
	case PawnList:
		if right, ok := rightValue.(PawnList); ok {
			result = left.v == right.v
		}
	case PawnFunctionPrimitive:
		if right, ok := rightValue.(PawnFunctionPrimitive); ok {
			result = left.v == right.v
		}
	case PawnFunctionUser:
		if right, ok := rightValue.(PawnFunctionUser); ok {
			result = left.v == right.v
		}
	default:
		panic(fmt.Sprint("Unexpected invalid type: ", leftValue))
	}

	return PawnBoolean{result}
}

func evaluateNotEqual(leftValue PawnValue, rightValue PawnValue) PawnValue {
	panic("TODO")
}

func evaluateGreaterThan(leftValue PawnValue, rightValue PawnValue) PawnValue {
		return numberOperation(
			leftValue, rightValue,
			func(a int64, b int64) PawnValue { return PawnBoolean{a > b} },
			func(a float64, b float64) PawnValue { return PawnBoolean{a > b} },
		)
}

// This function can return nil to indicate that the function that was called did
// not return anything.
func evaluateFunctionCall(expression ast.FunctionCall, environment Environment) PawnValue {
	callee := EvaluateExpression(expression.Function, environment)
	switch function := callee.(type) {
		case PawnFunctionUser:
			validateFunctionArguments(
				len(expression.PositionalArgs),
				len(function.v.positionalArguments),
				expression.NamedArgs,
				func(name string) bool {
					_, ok := function.v.namedArguments[name]
					return ok
				},
			)
			panic("TODO: Call the function")
		case PawnFunctionPrimitive:
			validateFunctionArguments(
				len(expression.PositionalArgs),
				int(function.v.positionalArgCount),
				expression.NamedArgs,
				func(name string) bool {
					return slices.Contains(function.v.namedArguments, name)
				},
			)
			panic("TODO: Call the function")
		default:
			panic(typeError("function", callee))
	}
}

func validateFunctionArguments(
	providedPositionalArgumentCount int,
	expectedPositionalArgumentCount int,
	providedNamedArgs []ast.NamedArg,
	validateNamedArg func(string) bool,
) {
	 if providedPositionalArgumentCount != expectedPositionalArgumentCount {
		panic(mismatchedArgumentCountError(
			expectedPositionalArgumentCount, providedPositionalArgumentCount,
		))
	 }

	 unexpectedNames := []string{}
	 for _, argument := range providedNamedArgs {
	 	if !validateNamedArg(argument.Name) {
			unexpectedNames = append(unexpectedNames, argument.Name)
	 	}
	 }
	 if len(unexpectedNames) != 0 {
		panic(unexpectedNamedArgumentsError(unexpectedNames))
	 }
}

func numberOperation(
	leftValue PawnValue,
	rightValue PawnValue,
	intHandler func(a int64, b int64) PawnValue,
	floatHandler func(a float64, b float64) PawnValue,
) PawnValue {
		switch a := leftValue.(type) {
			case PawnInt:
				switch b := rightValue.(type) {
					case PawnInt:
						return intHandler(a.v, b.v)
					case PawnFloat:
						return floatHandler(float64(a.v), b.v)
					default:
						panic(typeError("int or float", b))
				}
			case PawnFloat:
				switch b := rightValue.(type) {
					case PawnInt:
						return floatHandler(a.v, float64(b.v))
					case PawnFloat:
						return floatHandler(a.v, b.v)
					default:
						panic(typeError("int or float", b))
				}
			default:
				panic(typeError("int or float", a))
		}
}

func evaluateAnd(leftValue PawnValue, rightValue PawnValue) PawnBoolean {
	if leftBool, ok := leftValue.(PawnBoolean); ok {
		if rightBool, ok := rightValue.(PawnBoolean); ok {
			if leftBool.v && rightBool.v {
				return PawnBoolean{true}
			} else {
				return PawnBoolean{false}
			}
		} else {
			panic(typeError("bool", rightBool))
		}
		 
	} else {
		panic(typeError("bool", leftBool))
	}
}

func evaluateOr(leftValue PawnValue, rightValue PawnValue) PawnValue {
	if leftBool, ok := leftValue.(PawnBoolean); ok {
		if rightBool, ok := rightValue.(PawnBoolean); ok {
			if leftBool.v || rightBool.v {
				return PawnBoolean{true}
			} else {
				return PawnBoolean{false}
			}
		} else {
			panic(typeError("bool", rightBool))
		}
		 
	} else {
		panic(typeError("bool", leftBool))
	}

}
