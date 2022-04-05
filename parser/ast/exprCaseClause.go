package ast

type ExprCaseClause struct {
}

func (s ExprCaseClause) String() {
	//TODO implement me
	panic("implement me")
}

func (s ExprCaseClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ExprCaseClause)(nil)
