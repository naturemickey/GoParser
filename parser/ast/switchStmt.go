package ast

type SwitchStmt struct {
}

func (s SwitchStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s SwitchStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*SwitchStmt)(nil)
