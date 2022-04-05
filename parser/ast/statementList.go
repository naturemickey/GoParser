package ast

type StatementList struct {
	BaseNode
}

func (s *StatementList) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *StatementList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *StatementList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*StatementList)(nil)
