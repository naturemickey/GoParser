package ast

type EmptStmt struct {
	BaseNode
}

func (s *EmptStmt) codeBuilder() *CodeBuilder {
	return NewCodeBuilder()
}

func (s *EmptStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *EmptStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*EmptStmt)(nil)
