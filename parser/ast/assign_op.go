package ast

type AssignOp struct {
}

func (s AssignOp) String() {
	//TODO implement me
	panic("implement me")
}

func (s AssignOp) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*AssignOp)(nil)
