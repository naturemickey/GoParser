package ast

type LiteralValue struct {
}

func (s LiteralValue) String() {
	//TODO implement me
	panic("implement me")
}

func (s LiteralValue) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*LiteralValue)(nil)
