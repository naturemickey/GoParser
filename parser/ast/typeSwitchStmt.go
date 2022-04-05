package ast

type TypeSwitchStmt struct {
}

func (s TypeSwitchStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeSwitchStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeSwitchStmt)(nil)
