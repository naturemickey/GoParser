package ast

type TypeSwitchStmt struct {
	BaseNode
}

func (s *TypeSwitchStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeSwitchStmt)(nil)
