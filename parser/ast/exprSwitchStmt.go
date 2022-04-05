package ast

type ExprSwitchStmt struct {
	BaseNode
}

func (s *ExprSwitchStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ExprSwitchStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ExprSwitchStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ExprSwitchStmt)(nil)
