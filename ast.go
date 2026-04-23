package main

type AstNode struct {}

type Statement struct {
	AstNode
}

type If struct {
	Statement
	condition Expression
	body []Statement
	else_ []Statement
}

type While struct {
	Statement
	condition Expression
	body []Statement
}

type Apply struct {
	Statement
	subject Expression
	body []Statement
}

type FunctionDecl struct {
	Statement
	name string
	positionalArgs []string
	namedArgs []NamedArg
}

type NamedArg struct {
	Statement
	name string
	defaultValue Expression
}

type Return struct {
	Statement
	value *Expression
}

type Assignment struct {
	Statement
	destination Member
	value Expression
}

type ExpressionStatement struct {
	Statement
	expression Expression
}

type Expression struct {
	AstNode
}

type BinaryOp struct {
	Expression
	left Expression
	right Expression
	operator Operator
}

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
	Expression
	function Expression
	positionalArgs []Expression
	namedArgs []NamedArg
	body []Statement
}

type List struct {
	elements []Expression
}

type Literal[T any] struct {
	Expression
	Value T
}

type FloatLiteral Literal[float64]
type IntLiteral Literal[int64]
type BooleanLiteral Literal[bool]
type StringLiteral Literal[string]

type MemberExpression struct {
	Expression
	member Member
}

type Member struct {}

type Name struct {
	Member
	name string
}

type Property struct {
	Member
	object Expression
	property string
}

type Index struct {
	Member
	list Expression
	index []Expression
}
