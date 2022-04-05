package ast

type PrimaryExpr struct {
}

func (s PrimaryExpr) String() {
	//TODO implement me
	panic("implement me")
}

func (s PrimaryExpr) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*PrimaryExpr)(nil)
