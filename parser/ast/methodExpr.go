package ast

type MethodExpr struct {
}

func (s MethodExpr) String() {
	//TODO implement me
	panic("implement me")
}

func (s MethodExpr) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*MethodExpr)(nil)
