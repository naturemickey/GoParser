package ast

type ArrayLength struct {
	BaseNode
}

func (s *ArrayLength) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ArrayLength) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ArrayLength) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ArrayLength)(nil)
