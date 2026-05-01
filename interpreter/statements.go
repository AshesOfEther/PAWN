package interpreter

import (
	"fmt"

	"pawn/ast"
)

func EvaluateStatement(statement ast.Statement, environment Environment) {
	switch statement := statement.(type) {
		case ast.If:
			evaluateIf(statement, environment)
		case ast.While:
			panic("TODO")
		case ast.Apply:
			panic("TODO")
		case ast.FunctionDecl:
			panic("TODO")
		case ast.Return:
			panic("TODO")
		case ast.Assignment:
			panic("TODO")
		case ast.ExpressionStatement:
			panic("TODO")
		default:
			panic(fmt.Sprintf("Unexpected invalid Statement: %t", statement))
	}
}

func evaluateIf(statement ast.If, environment Environment) {
	if evaluateIfClause(statement.First, environment) {
		return
	}
	
	for _, clause := range statement.Rest {
		if evaluateIfClause(clause, environment) {
			return
		}
	}

	innerEnvironment := NewEnvironment(&environment)
	EvaluateStatements(statement.Else_, innerEnvironment)
	
}

func evaluateIfClause(clause ast.IfClause, environment Environment) bool {
	conditionValue := EvaluateExpression(clause.Condition, environment)
	if booleanValue, ok := conditionValue.(PawnBoolean); ok {
		if booleanValue.v {
			innerEnvironment := NewEnvironment(&environment)
			EvaluateStatements(clause.Body, innerEnvironment)
		}
		return booleanValue.v
	} else {
		panic(typeError("boolean", conditionValue))
	}
}

func EvaluateStatements(statements []ast.Statement, environment Environment) {
	for _, statement := range statements {
		EvaluateStatement(statement, environment)
	}
}
