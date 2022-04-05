package ast

type ExprSwitchCase struct {
}

func (s ExprSwitchCase) String() {
	//TODO implement me
	panic("implement me")
}

func (s ExprSwitchCase) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ExprSwitchCase)(nil)
