package ast

type SelectStmt struct {
	BaseNode
}

func (s *SelectStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *SelectStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SelectStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*SelectStmt)(nil)
