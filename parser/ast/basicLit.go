package ast

type BasicLit struct {
}

func (s BasicLit) String() {
	//TODO implement me
	panic("implement me")
}

func (s BasicLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*BasicLit)(nil)
