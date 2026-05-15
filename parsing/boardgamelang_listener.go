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

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterMemberExpr is called when entering the MemberExpr production.
	EnterMemberExpr(c *MemberExprContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterComparison is called when entering the Comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterLiteralExpr is called when entering the LiteralExpr production.
	EnterLiteralExpr(c *LiteralExprContext)

	// EnterIndex is called when entering the Index production.
	EnterIndex(c *IndexContext)

	// EnterRange is called when entering the Range production.
	EnterRange(c *RangeContext)

	// EnterPower is called when entering the Power production.
	EnterPower(c *PowerContext)

	// EnterProperty is called when entering the Property production.
	EnterProperty(c *PropertyContext)

	// EnterName is called when entering the Name production.
	EnterName(c *NameContext)

	// EnterList is called when entering the List production.
	EnterList(c *ListContext)

	// EnterBool is called when entering the Bool production.
	EnterBool(c *BoolContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

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

	// EnterApply is called when entering the Apply production.
	EnterApply(c *ApplyContext)

	// EnterWhile is called when entering the While production.
	EnterWhile(c *WhileContext)

	// EnterIf is called when entering the If production.
	EnterIf(c *IfContext)

	// EnterFunctionDecl is called when entering the FunctionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterReturn is called when entering the Return production.
	EnterReturn(c *ReturnContext)

	// EnterExpressionStatement is called when entering the ExpressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterAssignment is called when entering the Assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterFunctionDeclArgs is called when entering the functionDeclArgs production.
	EnterFunctionDeclArgs(c *FunctionDeclArgsContext)

	// EnterNameList is called when entering the nameList production.
	EnterNameList(c *NameListContext)

	// EnterNamedArgsDeclList is called when entering the namedArgsDeclList production.
	EnterNamedArgsDeclList(c *NamedArgsDeclListContext)

	// EnterNamedArgDecl is called when entering the namedArgDecl production.
	EnterNamedArgDecl(c *NamedArgDeclContext)

	// ExitCall is called when exiting the Call production.
	ExitCall(c *CallContext)

	// ExitMulDivMod is called when exiting the MulDivMod production.
	ExitMulDivMod(c *MulDivModContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitMemberExpr is called when exiting the MemberExpr production.
	ExitMemberExpr(c *MemberExprContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitComparison is called when exiting the Comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitLiteralExpr is called when exiting the LiteralExpr production.
	ExitLiteralExpr(c *LiteralExprContext)

	// ExitIndex is called when exiting the Index production.
	ExitIndex(c *IndexContext)

	// ExitRange is called when exiting the Range production.
	ExitRange(c *RangeContext)

	// ExitPower is called when exiting the Power production.
	ExitPower(c *PowerContext)

	// ExitProperty is called when exiting the Property production.
	ExitProperty(c *PropertyContext)

	// ExitName is called when exiting the Name production.
	ExitName(c *NameContext)

	// ExitList is called when exiting the List production.
	ExitList(c *ListContext)

	// ExitBool is called when exiting the Bool production.
	ExitBool(c *BoolContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

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

	// ExitApply is called when exiting the Apply production.
	ExitApply(c *ApplyContext)

	// ExitWhile is called when exiting the While production.
	ExitWhile(c *WhileContext)

	// ExitIf is called when exiting the If production.
	ExitIf(c *IfContext)

	// ExitFunctionDecl is called when exiting the FunctionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitReturn is called when exiting the Return production.
	ExitReturn(c *ReturnContext)

	// ExitExpressionStatement is called when exiting the ExpressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitAssignment is called when exiting the Assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitFunctionDeclArgs is called when exiting the functionDeclArgs production.
	ExitFunctionDeclArgs(c *FunctionDeclArgsContext)

	// ExitNameList is called when exiting the nameList production.
	ExitNameList(c *NameListContext)

	// ExitNamedArgsDeclList is called when exiting the namedArgsDeclList production.
	ExitNamedArgsDeclList(c *NamedArgsDeclListContext)

	// ExitNamedArgDecl is called when exiting the namedArgDecl production.
	ExitNamedArgDecl(c *NamedArgDeclContext)
}
