package ast

type ImportSpec struct {
	BaseNode
}

func (s *ImportSpec) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ImportSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ImportSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ImportSpec)(nil)
