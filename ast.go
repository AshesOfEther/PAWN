package main

type Statement interface {
	statementNode()
}

type If struct {
	first IfClause
	rest []IfClause
	else_ []Statement
}

type IfClause struct {
	condition Expression
	body []Statement 
}

func (If) statementNode() {}

type While struct {
	condition Expression
	body []Statement
}

func (While) statementNode() {}

type Apply struct {
	subject Expression
	body []Statement
}

func (Apply) statementNode() {}

type FunctionDecl struct {
	name string
	positionalArgs []string
	namedArgs []NamedArg
	body []Statement
}

func (FunctionDecl) statementNode() {}

type NamedArg struct {
	name string
	defaultValue Expression
}

type Return struct {
	value *Expression
}

func (Return) statementNode() {}

type Assignment struct {
	destination Member
	value Expression
}

func (Assignment) statementNode() {}

type ExpressionStatement struct {
	expression Expression
}

func (ExpressionStatement) statementNode() {}

type Expression interface {
	expressionNode()
}

type BinaryOp struct {
	left Expression
	right Expression
	operator Operator
}

func (BinaryOp) expressionNode() {}

type Operator int

const (
	Equal Operator = iota
	NotEqual
	LessThan
	GreaterThan
	LessOrEqual
	GreaterOrEqual
	Range
	RangeInclusive
	Add
	Subtract
	Multiply
	Divide
	Modulo
	Power
	And
	Or
)

type FunctionCall struct {
	function Expression
	positionalArgs []Expression
	namedArgs []NamedArg
	body []Statement
}

func (FunctionCall) expressionNode() {}

type List struct {
	elements []Expression
}

func (List) expressionNode() {}

type Literal[T any] struct {
	Value T
}

type FloatLiteral Literal[float64]
type IntLiteral Literal[int64]
type BooleanLiteral Literal[bool]
type StringLiteral Literal[string]

func (FloatLiteral) expressionNode() {}
func (IntLiteral) expressionNode() {}
func (BooleanLiteral) expressionNode() {}
func (StringLiteral) expressionNode() {}

type MemberExpression struct {
	member Member
}

func (MemberExpression) expressionNode() {}

type Member interface {
	memberNode()
}

type Name struct {
	name string
}

func (Name) memberNode() {}

type Property struct {
	object Expression
	property string
}

func (Property) memberNode() {}

type Index struct {
	list Expression
	index []Expression
}

func (Index) memberNode() {}
