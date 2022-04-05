package ast

type CommClause struct {
	BaseNode
}

func (s *CommClause) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *CommClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *CommClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*CommClause)(nil)
