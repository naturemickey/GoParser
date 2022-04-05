package ast

type RecvStmt struct {
}

func (s RecvStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s RecvStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*RecvStmt)(nil)
