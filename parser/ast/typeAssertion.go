package ast

type TypeAssertion struct {
	BaseNode

	type_ *Type_
}

func NewTypeAssertion(type_ *Type_) *TypeAssertion {
	return &TypeAssertion{type_: type_}
}

func (s *TypeAssertion) Type_() *Type_ {
	return s.type_
}

func (s *TypeAssertion) SetType_(type_ *Type_) {
	s.type_ = type_
}

func (s *TypeAssertion) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString(".(").appendNode(s.type_).appendString(")")
}

func (s *TypeAssertion) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeAssertion) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeAssertion)(nil)
