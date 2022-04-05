package ast

type Integer struct {
	BaseNode
}

func (s *Integer) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Integer) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Integer) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Integer)(nil)
