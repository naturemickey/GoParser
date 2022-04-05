package ast

type SliceType struct {
}

func (s SliceType) String() {
	//TODO implement me
	panic("implement me")
}

func (s SliceType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*SliceType)(nil)
