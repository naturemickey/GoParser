package ast

type LabeledStmt struct {
}

func (s LabeledStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s LabeledStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*LabeledStmt)(nil)
