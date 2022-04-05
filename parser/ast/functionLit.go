package ast

type FunctionLit struct {
	BaseNode
}

func (s *FunctionLit) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionLit) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*FunctionLit)(nil)
