package ast

type TypeList struct {
	BaseNode
}

func (s *TypeList) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeList)(nil)
