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
			evaluateWhile(statement, environment)
		case ast.Apply:
			panic("TODO")
		case ast.FunctionDecl:
			evaluateFunctionDeclaration(statement, environment)
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

func evaluateFunctionDeclaration(statement ast.FunctionDecl, environment Environment) {
	namedArgsMap := map[string]ast.Expression{}
	for _, namedArg := range statement.NamedArgs {
		if _, ok := namedArgsMap[namedArg.Name]; ok {
			panic(duplicateNamedArgumentDeclarationError(namedArg.Name))
		}
		namedArgsMap[namedArg.Name] = namedArg.DefaultValue
	}
	environment.variables[statement.Name] = PawnFunctionUser{
		&PawnFunctionUserInner{
			environment,
			statement.PositionalArgs,
			namedArgsMap,
			statement.Body,
		},
	}
}

func EvaluateStatements(statements []ast.Statement, environment Environment) {
	for _, statement := range statements {
		EvaluateStatement(statement, environment)
	}
}

func evaluateWhile(whileStmt ast.While, environment Environment) {
	for {
		conditionValue := EvaluateExpression(whileStmt.Condition, environment)
		booleanValue, ok := conditionValue.(PawnBoolean)
		if !ok {
			panic(typeError("boolean", conditionValue))
		}
		if booleanValue.v {
			EvaluateStatements(whileStmt.Body, NewEnvironment(&environment))
		} else {
			return
		}
	}
}
