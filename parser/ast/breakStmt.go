package ast

type BreakStmt struct {
	BaseNode
}

func (s *BreakStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *BreakStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *BreakStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *BreakStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*BreakStmt)(nil)
