package ast

type ContinueStmt struct {
	BaseNode
}

func (s *ContinueStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *ContinueStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ContinueStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ContinueStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*ContinueStmt)(nil)
