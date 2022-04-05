package ast

type ElementList struct {
}

func (s ElementList) String() {
	//TODO implement me
	panic("implement me")
}

func (s ElementList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ElementList)(nil)
