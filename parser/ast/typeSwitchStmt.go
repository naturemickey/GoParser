package ast

type TypeSwitchStmt struct {
	BaseNode
}

func (s *TypeSwitchStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchStmt) _SwitchStmt_() {
	//TODO implement me
	panic("implement me")
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

var _ SwitchStmt = (*TypeSwitchStmt)(nil)
