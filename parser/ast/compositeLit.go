package ast

type CompositeLit struct {
}

func (s CompositeLit) String() {
	//TODO implement me
	panic("implement me")
}

func (s CompositeLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*CompositeLit)(nil)
