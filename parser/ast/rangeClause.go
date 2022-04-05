package ast

type RangeClause struct {
}

func (s RangeClause) String() {
	//TODO implement me
	panic("implement me")
}

func (s RangeClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*RangeClause)(nil)
