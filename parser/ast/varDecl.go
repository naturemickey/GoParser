package ast

type VarDecl struct {
	BaseNode
}

func (s *VarDecl) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *VarDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *VarDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*VarDecl)(nil)
