package ast

type IncDecStmt struct {
}

func (s IncDecStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s IncDecStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*IncDecStmt)(nil)
