package main

import "fmt"

type Environment struct {
	variables map[string]PawnValue
	parent *Environment
}

func newEnvironment(parent *Environment) Environment {
	return Environment{
		map[string]PawnValue{},
		parent,
	}
}

// Expressions

func evaluateExpression(expression Expression, environment Environment) PawnValue {
	switch expression := expression.(type) {
		case BinaryOp:
			return evaluateBinaryOp(expression, environment)
		case FunctionCall:
			panic("TODO")
		case List:
			panic("TODO")
		case FloatLiteral:
			return PawnFloat{expression.Value}
		case IntLiteral:
			return PawnInt{expression.Value}
		case BooleanLiteral:
			return PawnBoolean{expression.Value}
		case StringLiteral:
			return PawnString{expression.Value}
		case MemberExpression:
			panic("TODO")
		default:
			panic(fmt.Sprintf("Unexpected invalid Expression: %t", expression))
	}
}

func evaluateBinaryOp(expression BinaryOp, environment Environment) PawnValue {
	leftValue := evaluateExpression(expression.left, environment)
	rightValue := evaluateExpression(expression.right, environment)
	
	switch expression.operator{
		case Equal:
			return evaluateEqual(leftValue, rightValue)
		case NotEqual:
			return evaluateNotEqual(leftValue, rightValue)
		case GreaterThan:
			return evaluateGreaterThan(leftValue, rightValue)
		case Add:
			panic("TODO")
		case And:
			panic("TODO")
		case Divide:
			panic("TODO")
		case GreaterOrEqual:
			panic("TODO")
		case LessOrEqual:
			panic("TODO")
		case LessThan:
			panic("TODO")
		case Modulo:
			panic("TODO")
		case Multiply:
			panic("TODO")
		case Or:
			panic("TODO")
		case Power:
			panic("TODO")
		case Range:
			panic("TODO")
		case RangeInclusive:
			panic("TODO")
		case Subtract:
			panic("TODO")
		default:
			panic(fmt.Sprintf("unexpected main.Operator: %#v", expression.operator))
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

// Statements

func evaluateStatement(statement Statement, environment Environment) {
	switch statement := statement.(type) {
		case If:
			evaluateIf(statement, environment)
		case While:
			panic("TODO")
		case Apply:
			panic("TODO")
		case FunctionDecl:
			panic("TODO")
		case Return:
			panic("TODO")
		case Assignment:
			panic("TODO")
		case ExpressionStatement:
			panic("TODO")
		default:
			panic(fmt.Sprintf("Unexpected invalid Statement: %t", statement))
	}
}

func evaluateIf(statement If, environment Environment) {
	if evaluateIfClause(statement.first, environment) {
		return
	}
	
	for _, clause := range statement.rest {
		if evaluateIfClause(clause, environment) {
			return
		}
	}

	innerEnvironment := newEnvironment(&environment)
	evaluateStatements(statement.else_, innerEnvironment)
	
}

func evaluateIfClause(clause IfClause, environment Environment) bool {
	conditionValue := evaluateExpression(clause.condition, environment)
	if booleanValue, ok := conditionValue.(PawnBoolean); ok {
		if booleanValue.v {
			innerEnvironment := newEnvironment(&environment)
			evaluateStatements(clause.body, innerEnvironment)
		}
		return booleanValue.v
	} else {
		panic(typeError("boolean", conditionValue))
	}
}

func evaluateStatements(statements []Statement, environment Environment) {
	for _, statement := range statements {
		evaluateStatement(statement, environment)
	}
}

type PawnError struct {
	message string
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
