package ast

type TypeSpec struct {
	BaseNode
}

func (s *TypeSpec) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeSpec)(nil)
