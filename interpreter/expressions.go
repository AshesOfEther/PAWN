package interpreter

import (
	"fmt"
	"math"
	"math/big"

	"pawn/ast"
)

func EvaluateExpression(expression ast.Expression, environment Environment) PawnValue {
	switch expression := expression.(type) {
		case ast.BinaryOp:
			return evaluateBinaryOp(expression, environment)
		case ast.FunctionCall:
			panic("TODO")
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
			return evaluateAdd(leftValue, rightValue)
		case ast.And:
			panic("TODO")
		case ast.Divide:
			return evaluateDivide(leftValue, rightValue)
		case ast.GreaterOrEqual:
			panic("TODO")
		case ast.LessOrEqual:
			panic("TODO")
		case ast.LessThan:
			panic("TODO")
		case ast.Modulo:
			return evaluateModulo(leftValue, rightValue)
		case ast.Multiply:
			return evaluateMultiply(leftValue, rightValue)
		case ast.Or:
			panic("TODO")
		case ast.Power:
			return evaluatePower(leftValue, rightValue)
		case ast.Range:
			panic("TODO")
		case ast.RangeInclusive:
			panic("TODO")
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
