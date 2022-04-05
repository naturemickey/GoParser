package ast

type ElementList struct {
	BaseNode
}

func (s *ElementList) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ElementList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ElementList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ElementList)(nil)
