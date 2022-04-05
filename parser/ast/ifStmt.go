package ast

type IfStmt struct {
}

func (s IfStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s IfStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*IfStmt)(nil)
