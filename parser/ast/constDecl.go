package ast

type ConstDecl struct {
}

func (s ConstDecl) String() {
	//TODO implement me
	panic("implement me")
}

func (s ConstDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ConstDecl)(nil)
