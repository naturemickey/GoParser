package ast

type ReturnStmt struct {
}

func (s ReturnStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s ReturnStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ReturnStmt)(nil)
