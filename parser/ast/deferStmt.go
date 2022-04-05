package ast

type DeferStmt struct {
}

func (s DeferStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s DeferStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*DeferStmt)(nil)
