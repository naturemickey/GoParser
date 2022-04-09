package ast

type Operand struct {
	BaseNode

	literal     Literal
	operandName string
	expression  *Expression
}

func NewOperand(literal Literal, operandName string, expression *Expression) *Operand {
	return &Operand{literal: literal, operandName: operandName, expression: expression}
}

func (s *Operand) Literal() Literal {
	return s.literal
}

func (s *Operand) SetLiteral(literal Literal) {
	s.literal = literal
}

func (s *Operand) OperandName() string {
	return s.operandName
}

func (s *Operand) SetOperandName(operandName string) {
	s.operandName = operandName
}

func (s *Operand) Expression() *Expression {
	return s.expression
}

func (s *Operand) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *Operand) codeBuilder() *CodeBuilder {
	if s.literal != nil {
		return s.literal.codeBuilder()
	}
	if s.operandName != "" {
		return NewCodeBuilder().appendString(s.operandName)
	}
	return NewCodeBuilder().appendString("(").appendNode(s.expression).appendString(")")
}

func (s *Operand) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Operand) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Operand)(nil)
