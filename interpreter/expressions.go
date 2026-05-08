package interpreter

import (
	"fmt"
	"math"
	"math/big"
	"reflect"

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
			return evaluateMemberExpression(expression, environment)
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
	result := false

	switch left := leftValue.(type) {
	case PawnBoolean:
		if right, ok := rightValue.(PawnBoolean); ok {
			result = left.v == right.v
		}
	case PawnObject:
		if right, ok := rightValue.(PawnObject); ok {
			result = reflect.ValueOf(left) == reflect.ValueOf(right)
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
			result = left == right
		}
	case PawnFunctionUser:
		if right, ok := rightValue.(PawnFunctionUser); ok {
			result = left == right
		}
	default:
		panic(fmt.Sprint("Unexpected invalid type: ", leftValue))
	}

	return PawnBoolean{result}
}

func evaluateNotEqual(leftValue PawnValue, rightValue PawnValue) PawnValue {
	return PawnBoolean{!evaluateEqual(leftValue, rightValue).(PawnBoolean).v}
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
					for _, namedArgument := range function.v.namedArguments {
						if namedArgument.name == name {
							return true
						}
					}
					return false
				},
			)
			if len(expression.Body) != 0 {
				panic(unexpectedBodyOnUserFunctionCall())
			}
			functionEnvironment := NewEnvironment(&function.v.environment)
			for i, name := range function.v.positionalArguments {
				functionEnvironment.variables[name] = EvaluateExpression(expression.PositionalArgs[i], environment)
			}
			for _, namedArgument := range expression.NamedArgs {
				functionEnvironment.variables[namedArgument.Name] = EvaluateExpression(namedArgument.DefaultValue, environment)
			}

			controlFlow := EvaluateStatements(function.v.body, functionEnvironment)
			if return_, ok := controlFlow.(Return); ok {
				return return_.value
			}
			return nil
		case PawnFunctionPrimitive:
			validateFunctionArguments(
				len(expression.PositionalArgs),
				len(function.v.positionalArgTypes),
				expression.NamedArgs,
				func(name string) bool {
					if function.v.namedArgumentTypes == nil {
						return false
					}
					_, ok := function.v.namedArgumentTypes[name]
					return ok
				},
			)
			if len(expression.Body) != 0 && !function.v.acceptsBody {
				panic(unexpectedBodyError())
			}
			
			arguments := []reflect.Value{}
			for i, positionalArgument := range expression.PositionalArgs {
				value := EvaluateExpression(positionalArgument, environment)
				reflectValue := reflect.ValueOf(value)
				expectedType := function.v.positionalArgTypes[i]
				if !reflectValue.Type().AssignableTo(expectedType) {
					panic(typeError(fmt.Sprintf("%s for positional argument #%d", getTypeName(expectedType), i), value))
				}
				arguments = append(arguments, reflectValue)
			}

			if function.v.namedArgumentStructType != nil {
				namedArguments := reflect.New(function.v.namedArgumentStructType).Elem()
				for _, namedArgument := range expression.NamedArgs {
					value := EvaluateExpression(namedArgument.DefaultValue, environment)
					reflectValue := reflect.ValueOf(value)
					field := namedArguments.FieldByName(toStructFieldCase(namedArgument.Name))
					if !reflectValue.Type().AssignableTo(field.Type().Elem()) {
						panic(typeError(fmt.Sprintf("%s for named argument '%s'", getTypeName(field.Type()), namedArgument.Name), value))
					}
					idk := reflect.New(field.Type().Elem())
					idk.Elem().Set(reflectValue)
					field.Set(idk)
				}
				arguments = append(arguments, namedArguments)
			}
			
			if function.v.acceptsBody {
				innerEnvironment := NewEnvironment(&environment)
				EvaluateStatements(expression.Body, innerEnvironment)
				body := innerEnvironment.variables
				arguments = append(arguments, reflect.ValueOf(body))
			}

			returnValue := function.v.f.Call(arguments)
			if len(returnValue) == 0 || returnValue[0].IsNil() {
				return nil
			}
			return returnValue[0].Interface().(PawnValue)
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

func evaluateMemberExpression(expression ast.MemberExpression, environment Environment) PawnValue {
	member := evaluateMember(expression.Member, environment)
	switch member := member.(type) {
		case PropertyOrVar:
			// It is impossible for the name to not exist since MemberReturnedMap only has a valid object and field
			return member.map_[member.key]
		case PawnList:
			return member
		case Index:
			// It is impossible to be out of bounds since MemberReturnedIndex only has a valid list and index
			return member.list[member.index]
		default:
			panic(fmt.Sprintf("Unexpected invalid Member return value: %t", member))
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
