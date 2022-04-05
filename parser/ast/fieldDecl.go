package ast

type FieldDecl struct {
	BaseNode
}

func (s *FieldDecl) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *FieldDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FieldDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*FieldDecl)(nil)
