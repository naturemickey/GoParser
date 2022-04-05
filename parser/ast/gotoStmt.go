package ast

type GotoStmt struct {
}

func (s GotoStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s GotoStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*GotoStmt)(nil)
