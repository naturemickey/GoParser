package ast

type TypeSpec struct {
	BaseNode
	id    string
	type_ *Type_
}

func (s *TypeSpec) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString(s.id).blank().appendNode(s.type_)
}

func (s *TypeSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeSpec)(nil)
