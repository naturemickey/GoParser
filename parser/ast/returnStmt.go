package ast

type ReturnStmt struct {
	BaseNode
}

func (s *ReturnStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ReturnStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ReturnStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ReturnStmt)(nil)
