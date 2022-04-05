package ast

type ContinueStmt struct {
}

func (s ContinueStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s ContinueStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ContinueStmt)(nil)
