package ast

type ReturnStmt struct {
	BaseNode
}

func (s *ReturnStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
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

var _ Statement = (*ReturnStmt)(nil)
