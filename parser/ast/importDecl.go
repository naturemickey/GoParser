package ast

type ImportDecl struct {
	BaseNode
}

func (s *ImportDecl) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ImportDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ImportDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ImportDecl)(nil)
