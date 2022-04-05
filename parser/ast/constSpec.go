package ast

type ConstSpec struct {
}

func (s ConstSpec) String() {
	//TODO implement me
	panic("implement me")
}

func (s ConstSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ConstSpec)(nil)
