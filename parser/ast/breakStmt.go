package ast

type BreakStmt struct {
}

func (s BreakStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s BreakStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*BreakStmt)(nil)
