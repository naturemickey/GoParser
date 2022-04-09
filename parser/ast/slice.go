package ast

type Slice struct {
	BaseNode

	expression1 *Expression
	expression2 *Expression
	expression3 *Expression
}

func NewSlice(expression1 *Expression, expression2 *Expression, expression3 *Expression) *Slice {
	return &Slice{expression1: expression1, expression2: expression2, expression3: expression3}
}

func (s *Slice) Expression1() *Expression {
	return s.expression1
}

func (s *Slice) SetExpression1(expression1 *Expression) {
	s.expression1 = expression1
}

func (s *Slice) Expression2() *Expression {
	return s.expression2
}

func (s *Slice) SetExpression2(expression2 *Expression) {
	s.expression2 = expression2
}

func (s *Slice) Expression3() *Expression {
	return s.expression3
}

func (s *Slice) SetExpression3(expression3 *Expression) {
	s.expression3 = expression3
}

func (s *Slice) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("[")
	if s.expression1 != nil {
		cb.appendNode(s.expression1)
	}
	cb.appendString(":")
	if s.expression2 != nil {
		cb.appendNode(s.expression2)
	}
	if s.expression3 != nil {
		cb.appendString(":")
		cb.appendNode(s.expression3)
	}
	cb.appendString("]")
	return cb
}

func (s *Slice) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Slice) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Slice)(nil)
