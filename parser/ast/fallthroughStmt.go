package ast

type FallThroughStmt struct {
}

func (s FallThroughStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s FallThroughStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*FallThroughStmt)(nil)
