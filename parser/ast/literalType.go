package ast

type LiteralType struct {
}

func (s LiteralType) String() {
	//TODO implement me
	panic("implement me")
}

func (s LiteralType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*LiteralType)(nil)
