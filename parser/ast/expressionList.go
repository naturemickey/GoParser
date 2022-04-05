package ast

type ExpressionList struct {
	BaseNode
	expressions []*Expression
}

func (s *ExpressionList) Expressions() []*Expression {
	return s.expressions
}

func (s *ExpressionList) SetExpressions(expressions []*Expression) {
	s.expressions = expressions
}

func (s *ExpressionList) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	for _, expression := range s.expressions {
		cb.appendNode(expression).appendString(", ")
	}
	cb.deleteLast()
	return cb
}

func (s *ExpressionList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ExpressionList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ExpressionList)(nil)
