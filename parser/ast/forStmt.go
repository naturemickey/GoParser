package ast

type ForStmt struct {
}

func (s ForStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s ForStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ForStmt)(nil)
