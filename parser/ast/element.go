package ast

type Element struct {
}

func (s Element) String() {
	//TODO implement me
	panic("implement me")
}

func (s Element) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Element)(nil)
