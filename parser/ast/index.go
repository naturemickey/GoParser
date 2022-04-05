package ast

type Index struct {
	BaseNode

	expression *Expression
}

func (s *Index) Expression() *Expression {
	return s.expression
}

func (s *Index) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *Index) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("[").appendNode(s.expression).appendString("]")
}

func (s *Index) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Index) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Index)(nil)
