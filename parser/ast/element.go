package ast

type Element struct {
	BaseNode
}

func (s *Element) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Element) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Element) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Element)(nil)
