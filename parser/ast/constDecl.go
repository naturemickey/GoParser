package ast

type ConstDecl struct {
	BaseNode
}

func (s *ConstDecl) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ConstDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ConstDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ConstDecl)(nil)
