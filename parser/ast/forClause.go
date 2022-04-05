package ast

type ForClause struct {
	BaseNode
}

func (s *ForClause) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ForClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ForClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ForClause)(nil)
