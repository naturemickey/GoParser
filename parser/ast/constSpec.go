package ast

type ConstSpec struct {
	BaseNode
}

func (s *ConstSpec) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ConstSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ConstSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ConstSpec)(nil)
