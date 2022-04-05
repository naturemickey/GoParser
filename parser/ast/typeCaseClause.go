package ast

type TypeCaseClause struct {
}

func (s TypeCaseClause) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeCaseClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeCaseClause)(nil)
