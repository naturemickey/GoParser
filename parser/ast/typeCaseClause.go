package ast

type TypeCaseClause struct {
	BaseNode
}

func (s *TypeCaseClause) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeCaseClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeCaseClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeCaseClause)(nil)
