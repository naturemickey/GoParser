package ast

type StatementList struct {
}

func (s StatementList) String() {
	//TODO implement me
	panic("implement me")
}

func (s StatementList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*StatementList)(nil)
