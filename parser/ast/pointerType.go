package ast

type PointerType struct {
}

func (s PointerType) String() {
	//TODO implement me
	panic("implement me")
}

func (s PointerType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*PointerType)(nil)
