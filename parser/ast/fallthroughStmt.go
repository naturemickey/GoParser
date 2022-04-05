package ast

type FallThroughStmt struct {
	BaseNode
}

func (s *FallThroughStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *FallThroughStmt) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("fallthrough")
}

func (s *FallThroughStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FallThroughStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*FallThroughStmt)(nil)
