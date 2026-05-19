// Code generated from ./BoardGameLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // BoardGameLang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BoardGameLangParser.
type BoardGameLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BoardGameLangParser#Call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#MulDivMod.
	VisitMulDivMod(ctx *MulDivModContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#MemberExpr.
	VisitMemberExpr(ctx *MemberExprContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#LiteralExpr.
	VisitLiteralExpr(ctx *LiteralExprContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Range.
	VisitRange(ctx *RangeContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#List.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#namedArgs.
	VisitNamedArgs(ctx *NamedArgsContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#namedArg.
	VisitNamedArg(ctx *NamedArgContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Apply.
	VisitApply(ctx *ApplyContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#While.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#If.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#FunctionDecl.
	VisitFunctionDecl(ctx *FunctionDeclContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#ExpressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#Assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#functionDeclArgs.
	VisitFunctionDeclArgs(ctx *FunctionDeclArgsContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#nameList.
	VisitNameList(ctx *NameListContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#namedArgsDeclList.
	VisitNamedArgsDeclList(ctx *NamedArgsDeclListContext) interface{}

	// Visit a parse tree produced by BoardGameLangParser#namedArgDecl.
	VisitNamedArgDecl(ctx *NamedArgDeclContext) interface{}
}
