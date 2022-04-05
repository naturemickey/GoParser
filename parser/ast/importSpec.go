package ast

type ImportSpec struct {
}

func (s ImportSpec) String() {
	//TODO implement me
	panic("implement me")
}

func (s ImportSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ImportSpec)(nil)
