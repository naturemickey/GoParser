package ast

type Index struct {
	BaseNode
}

func (s *Index) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Index) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Index) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Index)(nil)
