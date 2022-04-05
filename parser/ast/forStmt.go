package ast

type ForStmt struct {
	BaseNode
}

func (s *ForStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ForStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ForStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ForStmt)(nil)
