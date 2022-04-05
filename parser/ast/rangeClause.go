package ast

type RangeClause struct {
	BaseNode
}

func (s *RangeClause) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *RangeClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *RangeClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*RangeClause)(nil)
