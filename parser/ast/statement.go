package ast

type Statement struct {
}

func (s Statement) String() {
	//TODO implement me
	panic("implement me")
}

func (s Statement) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Statement)(nil)
