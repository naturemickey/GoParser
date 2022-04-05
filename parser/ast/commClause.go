package ast

type CommClause struct {
}

func (s CommClause) String() {
	//TODO implement me
	panic("implement me")
}

func (s CommClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*CommClause)(nil)
