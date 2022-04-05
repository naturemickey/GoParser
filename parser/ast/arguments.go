package ast

type Arguments struct {
	BaseNode
}

func (s *Arguments) String() string {
	return s.codeBuilder().String()
}

func (s *Arguments) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Arguments) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Arguments)(nil)
