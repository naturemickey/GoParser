package ast

type SimpleStmt struct {
}

func (s SimpleStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s SimpleStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*SimpleStmt)(nil)
