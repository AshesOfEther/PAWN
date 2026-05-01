package ast

import (
	"fmt"
	"strconv"
	"strings"

	"pawn/parsing"
)

type AstGeneratorVisitor struct {
	parsing.BoardGameLangVisitor
}

func generateExpression(ctx parsing.IExpressionContext) Expression {
	switch ctx := ctx.(type) {
	case *parsing.CallContext:
		argList := ctx.ArgList()

		var positionalArguments []Expression
		if argList.ExpressionList() != nil {
			expressions := argList.ExpressionList().AllExpression()
			positionalArguments = make([]Expression, len(expressions))
			for i, expression := range expressions {
				positionalArguments[i] = generateExpression(expression)
			}
		} else {
			positionalArguments = make([]Expression, 0)
		}

		originalNamedArgs := argList.NamedArgs().AllNamedArg()
		var namedArguments []NamedArg
		if argList.NamedArgs() != nil {
			namedArguments = make([]NamedArg, len(originalNamedArgs))
			for i, namedArg := range originalNamedArgs {
				namedArguments[i] = NamedArg{
					namedArg.NAME().GetText(),
					generateExpression(namedArg.Expression()),
				}
			}
		} else {
			namedArguments = make([]NamedArg, 0)
		}

		var body []Statement
		if ctx.StatementList() != nil {
			body = GenerateStatements(ctx.StatementList().AllStatement())
		} else {
			body = make([]Statement, 0)
		}

		return FunctionCall{
			generateExpression(ctx.Expression()),
			positionalArguments,
			namedArguments,
			body,
		}

	case *parsing.MulDivModContext:
		var operator Operator
		operatorToken := ctx.GetOp()
		switch operatorToken.GetText() {
		case "*":
			operator = Multiply
		case "/":
			operator = Divide
		case "%":
			operator = Modulo
		default:
			panic(fmt.Sprintf("Unexpected invalid MulDivMod operator at %v", ctx.GetSourceInterval()))
		}
		return BinaryOp{
			generateExpression(ctx.Expression(0)),
			generateExpression(ctx.Expression(1)),
			operator,
		}

	case *parsing.MemberExprContext:
		return MemberExpression{
			generateMember(ctx.Member()),
		}

	case *parsing.AddSubContext:
		var operator Operator
		operatorToken := ctx.GetOp()
		switch operatorToken.GetText() {
		case "+":
			operator = Add
		case "-":
			operator = Subtract
		default:
			panic(fmt.Sprintf("Unexpected invalid AddSub operator at %v", ctx.GetSourceInterval()))
		}
		return BinaryOp{
			generateExpression(ctx.Expression(0)),
			generateExpression(ctx.Expression(1)),
			operator,
		}

	case *parsing.ComparisonContext:
		var operator Operator
		operatorToken := ctx.GetOp()
		switch operatorToken.GetText() {
		case "==":
			operator = Equal
		case "!=":
			operator = NotEqual
		default:
			panic(fmt.Sprintf("Unexpected invalid Comparison operator at %v", ctx.GetSourceInterval()))
		}
		return BinaryOp{
			generateExpression(ctx.Expression(0)),
			generateExpression(ctx.Expression(1)),
			operator,
		}

	case *parsing.ParensContext:
		return generateExpression(ctx.Expression())

	case *parsing.LiteralExprContext:
		switch literal := ctx.Literal().(type) {
		case *parsing.ListContext:
			return List{generateExpressions(literal.AllExpression())}

		case *parsing.BoolContext:
			return BooleanLiteral{literal.GetText() == "true"}

		case *parsing.NumberContext:
			number := literal.GetText()
			if strings.ContainsRune(number, '.') {
				float, _ := strconv.ParseFloat(number, 64)
				return FloatLiteral{float}
			} else {
				int, _ := strconv.ParseInt(number, 10, 64)
				return IntLiteral{int}
			}

		case *parsing.StringContext:
			string := literal.GetText()
			return StringLiteral{string[1 : len(string)-1]}
		}

	case *parsing.IndexContext:
		return MemberExpression{Index{
			generateExpression(ctx.Expression()),
			generateExpressions(ctx.ExpressionList().AllExpression()),
		}}

	case *parsing.RangeContext:
		operator := RangeInclusive
		if ctx.GetOp().GetText() == ".." {
			operator = Range
		}
		return BinaryOp{
			generateExpression(ctx.Expression(0)),
			generateExpression(ctx.Expression(1)),
			operator,
		}

	case *parsing.PowerContext:
		return BinaryOp{
			generateExpression(ctx.Expression(0)),
			generateExpression(ctx.Expression(1)),
			Power,
		}
	}
	panic(fmt.Sprintf("[Internal error] Unexpected expression type: %#v", ctx))
}

func generateExpressions(list []parsing.IExpressionContext) []Expression {
	result := make([]Expression, len(list))
	for i, expression := range list {
		result[i] = generateExpression(expression)
	}
	return result
}

func generateMember(ctx parsing.IMemberContext) Member {
	switch ctx := ctx.(type) {
	case *parsing.PropertyContext:
		return Property{
			MemberExpression{generateMember(ctx.Member())},
			ctx.NAME().GetText(),
		}
	case *parsing.NameContext:
		return Name{
			ctx.GetText(),
		}
	}
	panic(fmt.Sprintf("[Internal error] Unexpected member type: %#v", ctx))
}

func generateStatement(ctx parsing.IStatementContext) Statement {
	switch ctx := ctx.(type) {
	case *parsing.ApplyContext:
		return Apply{
			generateExpression(ctx.Expression()),
			GenerateStatements(ctx.StatementList().AllStatement()),
		}

	case *parsing.WhileContext:
		return While{
			generateExpression(ctx.Expression()),
			GenerateStatements(ctx.StatementList().AllStatement()),
		}

	case *parsing.IfContext:
		conditions := ctx.AllExpression()
		bodies := ctx.AllStatementList()

		first := IfClause{
			generateExpression(conditions[0]),
			GenerateStatements(bodies[0].AllStatement()),
		}

		rest := make([]IfClause, len(conditions)-1)
		for i, condition := range conditions[1:] {
			rest[i] = IfClause{
				generateExpression(condition),
				GenerateStatements(bodies[i+1].AllStatement()),
			}
		}

		var else_ []Statement
		if len(bodies) > len(conditions) {
			else_ = GenerateStatements(bodies[len(bodies)-1].AllStatement())
		} else {
			else_ = make([]Statement, 0)
		}

		return If{
			first,
			rest,
			else_,
		}

	case *parsing.FunctionDeclContext:
		arguments := ctx.FunctionDeclArgs()

		var positionalArgs []string
		if arguments.NameList() != nil {
			positionalArgTokens := arguments.NameList().AllNAME()
			positionalArgs := make([]string, len(positionalArgTokens))
			for i, token := range positionalArgTokens {
				positionalArgs[i] = token.GetText()
			}
		} else {
			positionalArgs = make([]string, 0)
		}

		var namedArgs []NamedArg
		if arguments.NamedArgsDeclList() != nil {
			declarations := arguments.NamedArgsDeclList().AllNamedArgDecl()
			namedArgs := make([]NamedArg, len(declarations))
			for i, declaration := range declarations {
				namedArgs[i] = NamedArg{
					declaration.NAME().GetText(),
					generateExpression(declaration.Expression()),
				}
			}
		} else {
			namedArgs = make([]NamedArg, 0)
		}

		return FunctionDecl{
			ctx.NAME().GetText(),
			positionalArgs,
			namedArgs,
			GenerateStatements(ctx.StatementList().AllStatement()),
		}

	case *parsing.ReturnContext:
		var expression Expression = nil
		if ctx.Expression() != nil {
			expression = generateExpression(ctx.Expression())
		}
		return Return{
			&expression,
		}

	case *parsing.ExpressionStatementContext:
		return ExpressionStatement{
			generateExpression(ctx.Expression()),
		}

	case *parsing.AssignmentContext:
		return Assignment{
			Name{
				ctx.NAME().GetText(),
			},
			generateExpression(ctx.Expression()),
		}
	}
	panic(fmt.Sprintf("[Internal error] Unexpected statement type: %#v", ctx))
}

func GenerateStatements(list []parsing.IStatementContext) []Statement {
	result := make([]Statement, len(list))
	for i, statement := range list {
		result[i] = generateStatement(statement)
	}
	return result
}
