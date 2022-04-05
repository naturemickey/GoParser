package ast

type TypeAssertion struct {
	BaseNode
}

func (s *TypeAssertion) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeAssertion) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeAssertion) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeAssertion)(nil)
