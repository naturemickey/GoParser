package ast

type IncDecStmt struct {
	BaseNode
}

func (s *IncDecStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *IncDecStmt) _SimpleStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *IncDecStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *IncDecStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *IncDecStmt) String() string {
	return s.codeBuilder().String()
}

var _ SimpleStmt = (*IncDecStmt)(nil)
