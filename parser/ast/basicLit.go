package ast

type BasicLit struct {
	BaseNode
}

func (s *BasicLit) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *BasicLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *BasicLit) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*BasicLit)(nil)
