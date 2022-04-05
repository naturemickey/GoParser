package ast

type ExpressionStmt struct {
}

func (s ExpressionStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s ExpressionStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ExpressionStmt)(nil)
