package ast

type ExprCaseClause struct {
	BaseNode
}

func (s *ExprCaseClause) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ExprCaseClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ExprCaseClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ExprCaseClause)(nil)
