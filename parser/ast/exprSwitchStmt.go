package ast

type ExprSwitchStmt struct {
	BaseNode
}

func (s *ExprSwitchStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *ExprSwitchStmt) _SwitchStmt_() {
	//TODO implement me
	panic("implement me")
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

var _ SwitchStmt = (*ExprSwitchStmt)(nil)
