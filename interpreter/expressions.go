package interpreter

import (
	"fmt"
	"math"
	"math/big"
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
			return evaluateList(expression, environment)
		case ast.FloatLiteral:
			return PawnFloat{expression.Value}
		case ast.IntLiteral:
			return PawnInt{expression.Value}
		case ast.BooleanLiteral:
			return PawnBoolean{expression.Value}
		case ast.StringLiteral:
			return PawnString{expression.Value}
		case ast.MemberExpression:
			member := EvaluateMember(expression.Member, environment)
			switch member := member.(type) {
				case MemberReturnedMap:
					return member.mutateMe[member.nameToMutateAt]
				case MemberReturnedList:
					return member.list
				case MemberReturnedIndex:
					if int64(len(*member.list)) < member.indexToMutateAt || 0 > member.indexToMutateAt {
						panic(PawnError{fmt.Sprint("Index out of bounds: ", member.indexToMutateAt)})
					}
					return (*member.list)[member.indexToMutateAt]
				default:
					panic(fmt.Sprintf("Unexpected invalid Member return value: %t", member))
			}
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
			return evaluateAdd(leftValue, rightValue)
		case ast.And:
			return evaluateAnd(leftValue, rightValue) 
		case ast.Divide:
			return evaluateDivide(leftValue, rightValue)
		case ast.GreaterOrEqual:
			return evaluateGreaterOrEqual(leftValue, rightValue)
		case ast.LessOrEqual:
			return evaluateLessOrEqual(leftValue, rightValue)
		case ast.LessThan:
			return evaluateLessThan(leftValue, rightValue)
		case ast.Modulo:
			return evaluateModulo(leftValue, rightValue)
		case ast.Multiply:
			return evaluateMultiply(leftValue, rightValue)
		case ast.Or:
			return evaluateOr(leftValue, rightValue)
		case ast.Power:
			return evaluatePower(leftValue, rightValue)
		case ast.Range:
			return evaluateRange(leftValue, rightValue)
		case ast.RangeInclusive:
			return evaluateRangeInclusive(leftValue, rightValue)
		case ast.Subtract:
			return evaluateSubtract(leftValue, rightValue)
		default:
			panic(fmt.Sprintf("unexpected ast.Operator: %#v", expression.Operator))
	}
}

func evaluateEqual(leftValue PawnValue, rightValue PawnValue) PawnValue {
	return PawnBoolean{leftValue == rightValue}
}

func evaluateNotEqual(leftValue PawnValue, rightValue PawnValue) PawnValue {
	return PawnBoolean{leftValue != rightValue}
}

func evaluateGreaterThan(leftValue PawnValue, rightValue PawnValue) PawnValue {
		return numberOperation(
			leftValue, rightValue,
			func(a int64, b int64) PawnValue { return PawnBoolean{a > b} },
			func(a float64, b float64) PawnValue { return PawnBoolean{a > b} },
		)
}

func evaluateLessThan(leftValue PawnValue, rightValue PawnValue) PawnValue {
	return numberOperation(
		leftValue, rightValue,
		func(a int64, b int64) PawnValue { return PawnBoolean{a < b} },
		func(a float64, b float64) PawnValue { return PawnBoolean{a < b} },
	)
}

func evaluateLessOrEqual(leftValue PawnValue, rightValue PawnValue) PawnValue {
	return numberOperation(
		leftValue, rightValue,
		func(a int64, b int64) PawnValue { return PawnBoolean{a <= b} },
		func(a float64, b float64) PawnValue { return PawnBoolean{a <= b} },
	)
}

func evaluateGreaterOrEqual(leftValue PawnValue, rightValue PawnValue) PawnValue  {
	return numberOperation(
		leftValue, rightValue,
		func(a int64, b int64) PawnValue { return PawnBoolean{a >= b} },
		func(a float64, b float64) PawnValue { return PawnBoolean{a >= b} },
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

func evaluateAdd(value1 PawnValue, value2 PawnValue) PawnValue {
	return numberOperation(
		value1, value2,
		func(a int64, b int64) PawnValue { return PawnInt{a + b} },
		func(a float64, b float64) PawnValue { return PawnFloat{a + b} },
	)
}

func evaluateSubtract(value1 PawnValue, value2 PawnValue) PawnValue {
	return numberOperation(
		value1, value2,
		func(a int64, b int64) PawnValue { return PawnInt{a - b} },
		func(a float64, b float64) PawnValue { return PawnFloat{a - b} },
	)
}

func evaluateMultiply(value1 PawnValue, value2 PawnValue) PawnValue {
	return numberOperation(
		value1, value2,
		func(a int64, b int64) PawnValue { return PawnInt{a * b} },
		func(a float64, b float64) PawnValue { return PawnFloat{a * b} },
	)
}

func evaluateDivide(value1 PawnValue, value2 PawnValue) PawnValue {
	return numberOperation(
		value1, value2,
			func(a int64, b int64) PawnValue { 
				if b == 0 {
					panic(divideByZeroError())
				}
				return PawnInt{a / b}	
			},
		func(a float64, b float64) PawnValue { return PawnFloat{a / b} },
	)
}

func evaluateModulo(value1 PawnValue, value2 PawnValue) PawnValue {
	return numberOperation(
		value1, value2,
		func(a int64, b int64) PawnValue { return PawnInt{a % b} },
		func(a float64, b float64) PawnValue { return PawnFloat{math.Mod(a, b)} },
	)
}

// helper func for using int64
func powInt(x, y int64) int64 {
	if y < 0 {
		panic(negativePowerError(y))
	} else {
		return big.NewInt(0).Exp(big.NewInt(x), big.NewInt(y), nil).Int64()
	}	
}

func evaluatePower(value1 PawnValue, value2 PawnValue) PawnValue {
	return numberOperation(
		value1, value2,
		func(a int64, b int64) PawnValue { return PawnInt{powInt(a, b)} },
		func(a float64, b float64) PawnValue { return PawnFloat{math.Pow(a, b)} },
	)
}

func evaluateRange(leftValue PawnValue, rightValue PawnValue) PawnValue {
	return numberOperation(
		leftValue, rightValue,
         func(a, b int64) PawnValue {
			if b <= a {
				return PawnList{&[]PawnValue{}} 
			}

			makeList := make([]PawnValue, b-a)

			for i := a; i < b; i++ {
				makeList[i-a] = PawnInt{i}
			}

			return PawnList{&makeList}
		 },
		 func(a float64, b float64) PawnValue { panic(rangeFloatError()) },
	)
}

func evaluateRangeInclusive(leftValue PawnValue, rightValue PawnValue) PawnValue {
	return numberOperation(
		leftValue, rightValue,
		func(a int64, b int64) PawnValue {
			if b <= a {
				return PawnList{&[]PawnValue{PawnInt{b}}} 
			} 

			makeList := make([]PawnValue, (b-a)+1)

			for i := a; i <= b; i++ {
				makeList[i-a] = PawnInt{i}
			}

			return PawnList{&makeList}
		},
		func(a, b float64) PawnValue { panic(rangeFloatError()) },
	)
}

func evaluateList(list ast.List, environment Environment) PawnList {
	var listLength int = len(list.Elements)

	newList := make([]PawnValue, listLength)

	for i := range listLength {
		newList[i] = EvaluateExpression(list.Elements[i],environment)
	}

	return PawnList{&newList}
}
