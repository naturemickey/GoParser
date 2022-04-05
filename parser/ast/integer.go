package ast

type Integer struct {
}

func (s Integer) String() {
	//TODO implement me
	panic("implement me")
}

func (s Integer) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Integer)(nil)
