package ast

type GotoStmt struct {
	BaseNode
}

func (s *GotoStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *GotoStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *GotoStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*GotoStmt)(nil)
