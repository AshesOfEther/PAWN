// Code generated from ./BoardGameLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // BoardGameLang
import "github.com/antlr4-go/antlr/v4"

// BaseBoardGameLangListener is a complete listener for a parse tree produced by BoardGameLangParser.
type BaseBoardGameLangListener struct{}

var _ BoardGameLangListener = &BaseBoardGameLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBoardGameLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBoardGameLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBoardGameLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBoardGameLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCall is called when production Call is entered.
func (s *BaseBoardGameLangListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production Call is exited.
func (s *BaseBoardGameLangListener) ExitCall(ctx *CallContext) {}

// EnterMulDivMod is called when production MulDivMod is entered.
func (s *BaseBoardGameLangListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production MulDivMod is exited.
func (s *BaseBoardGameLangListener) ExitMulDivMod(ctx *MulDivModContext) {}

// EnterMemberExpr is called when production MemberExpr is entered.
func (s *BaseBoardGameLangListener) EnterMemberExpr(ctx *MemberExprContext) {}

// ExitMemberExpr is called when production MemberExpr is exited.
func (s *BaseBoardGameLangListener) ExitMemberExpr(ctx *MemberExprContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseBoardGameLangListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseBoardGameLangListener) ExitAddSub(ctx *AddSubContext) {}

// EnterComparison is called when production Comparison is entered.
func (s *BaseBoardGameLangListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production Comparison is exited.
func (s *BaseBoardGameLangListener) ExitComparison(ctx *ComparisonContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseBoardGameLangListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseBoardGameLangListener) ExitParens(ctx *ParensContext) {}

// EnterLiteralExpr is called when production LiteralExpr is entered.
func (s *BaseBoardGameLangListener) EnterLiteralExpr(ctx *LiteralExprContext) {}

// ExitLiteralExpr is called when production LiteralExpr is exited.
func (s *BaseBoardGameLangListener) ExitLiteralExpr(ctx *LiteralExprContext) {}

// EnterIndex is called when production Index is entered.
func (s *BaseBoardGameLangListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production Index is exited.
func (s *BaseBoardGameLangListener) ExitIndex(ctx *IndexContext) {}

// EnterRange is called when production Range is entered.
func (s *BaseBoardGameLangListener) EnterRange(ctx *RangeContext) {}

// ExitRange is called when production Range is exited.
func (s *BaseBoardGameLangListener) ExitRange(ctx *RangeContext) {}

// EnterPower is called when production Power is entered.
func (s *BaseBoardGameLangListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production Power is exited.
func (s *BaseBoardGameLangListener) ExitPower(ctx *PowerContext) {}

// EnterMember is called when production member is entered.
func (s *BaseBoardGameLangListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseBoardGameLangListener) ExitMember(ctx *MemberContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseBoardGameLangListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseBoardGameLangListener) ExitLiteral(ctx *LiteralContext) {}

// EnterList is called when production list is entered.
func (s *BaseBoardGameLangListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseBoardGameLangListener) ExitList(ctx *ListContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseBoardGameLangListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseBoardGameLangListener) ExitArgList(ctx *ArgListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseBoardGameLangListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseBoardGameLangListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterNamedArgs is called when production namedArgs is entered.
func (s *BaseBoardGameLangListener) EnterNamedArgs(ctx *NamedArgsContext) {}

// ExitNamedArgs is called when production namedArgs is exited.
func (s *BaseBoardGameLangListener) ExitNamedArgs(ctx *NamedArgsContext) {}

// EnterNamedArg is called when production namedArg is entered.
func (s *BaseBoardGameLangListener) EnterNamedArg(ctx *NamedArgContext) {}

// ExitNamedArg is called when production namedArg is exited.
func (s *BaseBoardGameLangListener) ExitNamedArg(ctx *NamedArgContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseBoardGameLangListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseBoardGameLangListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseBoardGameLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseBoardGameLangListener) ExitStatement(ctx *StatementContext) {}

// EnterApply is called when production apply is entered.
func (s *BaseBoardGameLangListener) EnterApply(ctx *ApplyContext) {}

// ExitApply is called when production apply is exited.
func (s *BaseBoardGameLangListener) ExitApply(ctx *ApplyContext) {}

// EnterWhile is called when production while is entered.
func (s *BaseBoardGameLangListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BaseBoardGameLangListener) ExitWhile(ctx *WhileContext) {}

// EnterIf is called when production if is entered.
func (s *BaseBoardGameLangListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BaseBoardGameLangListener) ExitIf(ctx *IfContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BaseBoardGameLangListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BaseBoardGameLangListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterFunctionDeclArgs is called when production functionDeclArgs is entered.
func (s *BaseBoardGameLangListener) EnterFunctionDeclArgs(ctx *FunctionDeclArgsContext) {}

// ExitFunctionDeclArgs is called when production functionDeclArgs is exited.
func (s *BaseBoardGameLangListener) ExitFunctionDeclArgs(ctx *FunctionDeclArgsContext) {}

// EnterNameList is called when production nameList is entered.
func (s *BaseBoardGameLangListener) EnterNameList(ctx *NameListContext) {}

// ExitNameList is called when production nameList is exited.
func (s *BaseBoardGameLangListener) ExitNameList(ctx *NameListContext) {}

// EnterNamedArgsDeclList is called when production namedArgsDeclList is entered.
func (s *BaseBoardGameLangListener) EnterNamedArgsDeclList(ctx *NamedArgsDeclListContext) {}

// ExitNamedArgsDeclList is called when production namedArgsDeclList is exited.
func (s *BaseBoardGameLangListener) ExitNamedArgsDeclList(ctx *NamedArgsDeclListContext) {}

// EnterNamedArgDecl is called when production namedArgDecl is entered.
func (s *BaseBoardGameLangListener) EnterNamedArgDecl(ctx *NamedArgDeclContext) {}

// ExitNamedArgDecl is called when production namedArgDecl is exited.
func (s *BaseBoardGameLangListener) ExitNamedArgDecl(ctx *NamedArgDeclContext) {}

// EnterReturn is called when production return is entered.
func (s *BaseBoardGameLangListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BaseBoardGameLangListener) ExitReturn(ctx *ReturnContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseBoardGameLangListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseBoardGameLangListener) ExitAssignment(ctx *AssignmentContext) {}
