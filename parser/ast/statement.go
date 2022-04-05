package ast

type Statement struct {
	BaseNode
}

func (s *Statement) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Statement) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Statement) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Statement)(nil)
