package ast

type Conversion struct {
	BaseNode

	nonNamedType *NonNamedType
	expression   *Expression
}

func NewConversion(nonNamedType *NonNamedType, expression *Expression) *Conversion {
	return &Conversion{nonNamedType: nonNamedType, expression: expression}
}

func (s *Conversion) NonNamedType() *NonNamedType {
	return s.nonNamedType
}

func (s *Conversion) SetNonNamedType(nonNamedType *NonNamedType) {
	s.nonNamedType = nonNamedType
}

func (s *Conversion) Expression() *Expression {
	return s.expression
}

func (s *Conversion) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *Conversion) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendNode(s.nonNamedType).appendString("(").appendNode(s.expression).appendString(")")
}

func (s *Conversion) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Conversion) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Conversion)(nil)
