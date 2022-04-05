package ast

type SimpleStmt struct {
	BaseNode
}

func (s *SimpleStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SimpleStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*SimpleStmt)(nil)
