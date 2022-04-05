package ast

type IfStmt struct {
	BaseNode
}

func (s *IfStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *IfStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *IfStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *IfStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*IfStmt)(nil)
