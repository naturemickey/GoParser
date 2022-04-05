package ast

type ElementType struct {
	BaseNode
}

func (s *ElementType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ElementType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ElementType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ElementType)(nil)
