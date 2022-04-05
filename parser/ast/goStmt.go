package ast

type GoStmt struct {
}

func (s GoStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s GoStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*GoStmt)(nil)
