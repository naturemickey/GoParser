package ast

type ExpressionList struct {
	BaseNode
}

func (s *ExpressionList) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ExpressionList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ExpressionList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ExpressionList)(nil)
