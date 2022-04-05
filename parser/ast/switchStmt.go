package ast

type SwitchStmt struct {
	BaseNode
}

func (s *SwitchStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *SwitchStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SwitchStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*SwitchStmt)(nil)
