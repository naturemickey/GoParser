package ast

type TypeDecl struct {
	BaseNode
}

func (s *TypeDecl) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeDecl)(nil)
