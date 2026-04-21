// Code generated from ./BoardGameLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // BoardGameLang
import "github.com/antlr4-go/antlr/v4"

// BoardGameLangListener is a complete listener for a parse tree produced by BoardGameLangParser.
type BoardGameLangListener interface {
	antlr.ParseTreeListener

	// EnterCall is called when entering the Call production.
	EnterCall(c *CallContext)

	// EnterMulDivMod is called when entering the MulDivMod production.
	EnterMulDivMod(c *MulDivModContext)

	// EnterMemberExpr is called when entering the MemberExpr production.
	EnterMemberExpr(c *MemberExprContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterComparison is called when entering the Comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterLiteralExpr is called when entering the LiteralExpr production.
	EnterLiteralExpr(c *LiteralExprContext)

	// EnterIndex is called when entering the Index production.
	EnterIndex(c *IndexContext)

	// EnterRange is called when entering the Range production.
	EnterRange(c *RangeContext)

	// EnterPower is called when entering the Power production.
	EnterPower(c *PowerContext)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterNamedArgs is called when entering the namedArgs production.
	EnterNamedArgs(c *NamedArgsContext)

	// EnterNamedArg is called when entering the namedArg production.
	EnterNamedArg(c *NamedArgContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterApply is called when entering the apply production.
	EnterApply(c *ApplyContext)

	// EnterWhile is called when entering the while production.
	EnterWhile(c *WhileContext)

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterFunctionDeclArgs is called when entering the functionDeclArgs production.
	EnterFunctionDeclArgs(c *FunctionDeclArgsContext)

	// EnterNameList is called when entering the nameList production.
	EnterNameList(c *NameListContext)

	// EnterNamedArgsDeclList is called when entering the namedArgsDeclList production.
	EnterNamedArgsDeclList(c *NamedArgsDeclListContext)

	// EnterNamedArgDecl is called when entering the namedArgDecl production.
	EnterNamedArgDecl(c *NamedArgDeclContext)

	// EnterReturn is called when entering the return production.
	EnterReturn(c *ReturnContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// ExitCall is called when exiting the Call production.
	ExitCall(c *CallContext)

	// ExitMulDivMod is called when exiting the MulDivMod production.
	ExitMulDivMod(c *MulDivModContext)

	// ExitMemberExpr is called when exiting the MemberExpr production.
	ExitMemberExpr(c *MemberExprContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitComparison is called when exiting the Comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitLiteralExpr is called when exiting the LiteralExpr production.
	ExitLiteralExpr(c *LiteralExprContext)

	// ExitIndex is called when exiting the Index production.
	ExitIndex(c *IndexContext)

	// ExitRange is called when exiting the Range production.
	ExitRange(c *RangeContext)

	// ExitPower is called when exiting the Power production.
	ExitPower(c *PowerContext)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitNamedArgs is called when exiting the namedArgs production.
	ExitNamedArgs(c *NamedArgsContext)

	// ExitNamedArg is called when exiting the namedArg production.
	ExitNamedArg(c *NamedArgContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitApply is called when exiting the apply production.
	ExitApply(c *ApplyContext)

	// ExitWhile is called when exiting the while production.
	ExitWhile(c *WhileContext)

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitFunctionDeclArgs is called when exiting the functionDeclArgs production.
	ExitFunctionDeclArgs(c *FunctionDeclArgsContext)

	// ExitNameList is called when exiting the nameList production.
	ExitNameList(c *NameListContext)

	// ExitNamedArgsDeclList is called when exiting the namedArgsDeclList production.
	ExitNamedArgsDeclList(c *NamedArgsDeclListContext)

	// ExitNamedArgDecl is called when exiting the namedArgDecl production.
	ExitNamedArgDecl(c *NamedArgDeclContext)

	// ExitReturn is called when exiting the return production.
	ExitReturn(c *ReturnContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)
}
