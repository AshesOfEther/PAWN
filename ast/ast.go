package ast

type Statement interface {
	statementNode()
}

type If struct {
	First IfClause
	Rest []IfClause
	Else_ []Statement
}

type IfClause struct {
	Condition Expression
	Body []Statement
}

func (If) statementNode() {}

type While struct {
	Condition Expression
	Body []Statement
}

func (While) statementNode() {}

type Apply struct {
	Subject Expression
	Body []Statement
}

func (Apply) statementNode() {}

type FunctionDecl struct {
	Name string
	PositionalArgs []string
	NamedArgs []NamedArg
	Body []Statement
}

func (FunctionDecl) statementNode() {}

type NamedArg struct {
	Name string
	DefaultValue Expression
}

type Return struct {
	Value *Expression // Can be nil
}

func (Return) statementNode() {}

type Assignment struct {
	Destination Member
	Value Expression
}

func (Assignment) statementNode() {}

type ExpressionStatement struct {
	Expression Expression
}

func (ExpressionStatement) statementNode() {}

type Expression interface {
	expressionNode()
}

type BinaryOp struct {
	Left Expression
	Right Expression
	Operator Operator
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
	Function Expression
	PositionalArgs []Expression
	NamedArgs []NamedArg
	Body []Statement
}

func (FunctionCall) expressionNode() {}

type List struct {
	Elements []Expression
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
	Member Member
}

func (MemberExpression) expressionNode() {}

type Member interface {
	memberNode()
}

type Name struct {
	Name string
}

func (Name) memberNode() {}

type Property struct {
	Object Expression
	Property string
}

func (Property) memberNode() {}

type Index struct {
	List Expression
	Index []Expression
}

func (Index) memberNode() {}
