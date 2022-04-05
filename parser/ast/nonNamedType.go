package ast

type NonNamedType struct {
	BaseNode
}

func (s *NonNamedType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *NonNamedType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *NonNamedType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*NonNamedType)(nil)
