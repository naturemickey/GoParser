package ast

type ExpressionList struct {
}

func (s ExpressionList) String() {
	//TODO implement me
	panic("implement me")
}

func (s ExpressionList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ExpressionList)(nil)
