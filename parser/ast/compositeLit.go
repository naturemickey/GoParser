package ast

type CompositeLit struct {
	BaseNode
}

func (s *CompositeLit) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *CompositeLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *CompositeLit) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*CompositeLit)(nil)
