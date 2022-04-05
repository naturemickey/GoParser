package ast

type TypeSpec struct {
}

func (s TypeSpec) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeSpec)(nil)
