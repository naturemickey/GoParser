package ast

type FallThroughStmt struct {
	BaseNode
}

func (s *FallThroughStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *FallThroughStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FallThroughStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*FallThroughStmt)(nil)
