package ast

type Arguments struct {
}

func (s Arguments) String() {
	//TODO implement me
	panic("implement me")
}

func (s Arguments) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Arguments)(nil)
