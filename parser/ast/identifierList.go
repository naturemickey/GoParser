package ast

type IdentifierList struct {
	BaseNode
}

func (s *IdentifierList) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *IdentifierList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *IdentifierList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*IdentifierList)(nil)
