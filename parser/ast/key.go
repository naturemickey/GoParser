package ast

type Key struct {
}

func (s Key) String() {
	//TODO implement me
	panic("implement me")
}

func (s Key) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Key)(nil)
