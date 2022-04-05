package ast

type VarDecl struct {
}

func (s VarDecl) String() {
	//TODO implement me
	panic("implement me")
}

func (s VarDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*VarDecl)(nil)
