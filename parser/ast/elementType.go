package ast

type ElementType struct {
}

func (s ElementType) String() {
	//TODO implement me
	panic("implement me")
}

func (s ElementType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ElementType)(nil)
