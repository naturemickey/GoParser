package ast

type TypeLit struct {
}

func (s TypeLit) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeLit)(nil)
