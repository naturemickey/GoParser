package ast

type FunctionType struct {
	BaseNode
}

func (s *FunctionType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*FunctionType)(nil)
