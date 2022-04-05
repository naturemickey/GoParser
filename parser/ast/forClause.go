package ast

type ForClause struct {
}

func (s ForClause) String() {
	//TODO implement me
	panic("implement me")
}

func (s ForClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ForClause)(nil)
