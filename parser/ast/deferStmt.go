package ast

type DeferStmt struct {
	BaseNode
}

func (s *DeferStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *DeferStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *DeferStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *DeferStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*DeferStmt)(nil)
