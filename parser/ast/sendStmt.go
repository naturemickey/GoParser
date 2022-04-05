package ast

type SendStmt struct {
}

func (s SendStmt) String() {
	//TODO implement me
	panic("implement me")
}

func (s SendStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*SendStmt)(nil)
