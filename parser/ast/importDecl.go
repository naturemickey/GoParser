package ast

type ImportDecl struct {
}

func (s ImportDecl) String() {
	//TODO implement me
	panic("implement me")
}

func (s ImportDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ImportDecl)(nil)
