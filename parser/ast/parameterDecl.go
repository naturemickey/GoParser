package ast

type ParameterDecl struct {
	BaseNode
}

func (s *ParameterDecl) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ParameterDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ParameterDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ParameterDecl)(nil)
