package ast

type Expression struct {
}

func (s Expression) String() {
	//TODO implement me
	panic("implement me")
}

func (s Expression) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Expression)(nil)
