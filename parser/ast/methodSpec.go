package ast

type MethodSpec struct {
}

func (s MethodSpec) String() {
	//TODO implement me
	panic("implement me")
}

func (s MethodSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*MethodSpec)(nil)
