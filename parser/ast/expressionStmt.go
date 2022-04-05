package ast

type ExpressionStmt struct {
	BaseNode
}

func (s *ExpressionStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ExpressionStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ExpressionStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ExpressionStmt)(nil)
