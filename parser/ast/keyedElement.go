package ast

type KeyedElement struct {
}

func (s KeyedElement) String() {
	//TODO implement me
	panic("implement me")
}

func (s KeyedElement) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*KeyedElement)(nil)
