package ast

type TypeLit struct {
	BaseNode
}

func (s *TypeLit) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeLit) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeLit)(nil)
