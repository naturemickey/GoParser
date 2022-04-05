package ast

type SelectStmt struct {
}

func (s SelectStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s SelectStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*SelectStmt)(nil)
