package interpreter

import (
	"fmt"

	"pawn/ast"
)

func EvaluateStatement(statement ast.Statement, environment Environment) ControlFlow {
	switch statement := statement.(type) {
		case ast.If:
			return evaluateIf(statement, environment)
		case ast.While:
			return evaluateWhile(statement, environment)
		case ast.Apply:
			panic("TODO")
		case ast.FunctionDecl:
			evaluateFunctionDeclaration(statement, environment)
		case ast.Return:
			return evaluateReturn(statement, environment)
		case ast.Assignment:
			evaluateAssignment(statement, environment)
		case ast.ExpressionStatement:
			EvaluateExpressionStatement(statement, environment)
		default:
			panic(fmt.Sprintf("Unexpected invalid Statement: %t", statement))
	}
	return nil
}

func evaluateAssignment(statement ast.Assignment, environment Environment) {
	value := EvaluateExpression(statement.Value, environment)
	switch member := statement.Destination.(type) {
		case ast.Name:
			result := evaluateMemberName(member, environment)
			if result != nil {
				result.map_[result.key] = value
			} else {
				environment.variables[member.Name] = value
			}
		case ast.Property:
			result := evaluateMemberProperty(member, environment)
			result.map_[result.key] = value
		case ast.Index:
			switch result := evaluateMemberIndex(member, environment).(type) {
				case Index:
					result.list[result.index] = value
				case PawnList:
					panic(assignToMultipleIndicesError())
				default:
					panic(fmt.Sprintf("Unexpected invalid ResolvedMember: %t", result))
			}
		default:
			panic(fmt.Sprintf("Unexpected invalid Member: %t", member))
	}
}

func evaluateIf(statement ast.If, environment Environment) ControlFlow {
	if executed, controlFlow := evaluateIfClause(statement.First, environment); executed {
		return controlFlow
	}

	for _, clause := range statement.Rest {
		if executed, controlFlow := evaluateIfClause(clause, environment); executed {
			return controlFlow
		}
	}

	innerEnvironment := NewEnvironment(&environment)
	return EvaluateStatements(statement.Else_, innerEnvironment)

}

func evaluateIfClause(clause ast.IfClause, environment Environment) (bool, ControlFlow) {
	conditionValue := EvaluateExpression(clause.Condition, environment)
	if booleanValue, ok := conditionValue.(PawnBoolean); ok {
		var controlFlow ControlFlow
		if booleanValue.v {
			innerEnvironment := NewEnvironment(&environment)
			controlFlow = EvaluateStatements(clause.Body, innerEnvironment)
		}
		return booleanValue.v, controlFlow
	} else {
		panic(typeError("boolean", conditionValue))
	}
}

func evaluateFunctionDeclaration(statement ast.FunctionDecl, environment Environment) {
	// Ensure that no two named arguments have the same name.
	existingNamedArguments := map[string]bool{}
	namedArguments := make([]NamedArgument, len(statement.NamedArgs))
	for i, namedArg := range statement.NamedArgs {
		if _, ok := existingNamedArguments[namedArg.Name]; ok {
			panic(duplicateNamedArgumentDeclarationError(namedArg.Name))
		}
		existingNamedArguments[namedArg.Name] = true
		namedArguments[i] = NamedArgument{namedArg.Name, namedArg.DefaultValue}
	}
	environment.variables[statement.Name] = PawnFunctionUser{
		&PawnFunctionUserInner{
			environment,
			statement.PositionalArgs,
			namedArguments,
			statement.Body,
		},
	}
}

func evaluateReturn(statement ast.Return, environment Environment) Return {
	var value PawnValue
	if statement.Value != nil {
		value = EvaluateExpression(*statement.Value, environment)
	}
	return Return{value}
}

func EvaluateExpressionStatement(statement ast.ExpressionStatement, environment Environment) PawnValue {
	// `EvaluateExpression` errors if the function doesn't return a value, because
	// it's used in cases wwhere a value is expected. However, it's fine when the
	// function call is an expression statement, because a value is not needed.
	if call, ok := statement.Expression.(ast.FunctionCall); ok {
		return evaluateFunctionCall(call, environment)
	}

	return EvaluateExpression(statement.Expression, environment)
}

func EvaluateStatements(statements []ast.Statement, environment Environment) ControlFlow {
	for _, statement := range statements {
		controlFlow := EvaluateStatement(statement, environment)
		if controlFlow != nil {
			return controlFlow
		}
	}
	return nil
}

func evaluateWhile(whileStmt ast.While, environment Environment) ControlFlow {
	for {
		conditionValue := EvaluateExpression(whileStmt.Condition, environment)
		booleanValue, ok := conditionValue.(PawnBoolean)
		if !ok {
			panic(typeError("boolean", conditionValue))
		}
		if booleanValue.v {
			controlFlow := EvaluateStatements(whileStmt.Body, NewEnvironment(&environment))
			if controlFlow != nil {
				return controlFlow
			}
		} else {
			return nil
		}
	}
}

// Represents a control flow operation that the caller should act on. `nil`
// indicates that no special handling is expected. A caller encountering a
// `ControlFlow` instance that it doesn't act on should return it immediately
// (e.g. a return statement inside an if statement).
type ControlFlow interface {
	controlFlow()
}

// Represents that a return statement was encountered.
type Return struct {
	value PawnValue
}

func (Return) controlFlow() {}
