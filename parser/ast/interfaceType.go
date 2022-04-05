package ast

type InterfaceType struct {
	BaseNode
}

func (s *InterfaceType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *InterfaceType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *InterfaceType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*InterfaceType)(nil)
