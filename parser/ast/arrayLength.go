package ast

type ArrayLength struct {
}

func (s ArrayLength) String() {
	//TODO implement me
	panic("implement me")
}

func (s ArrayLength) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ArrayLength)(nil)
