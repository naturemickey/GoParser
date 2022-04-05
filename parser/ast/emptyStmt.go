package ast

type EmptStmt struct {
}

func (s EmptStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s EmptStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*EmptStmt)(nil)
