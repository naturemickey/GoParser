package ast

type ExprSwitchStmt struct {
}

func (s ExprSwitchStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s ExprSwitchStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ExprSwitchStmt)(nil)
