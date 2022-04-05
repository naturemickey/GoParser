package ast

type ImportPath struct {
	BaseNode
}

func (s *ImportPath) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ImportPath) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ImportPath) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ImportPath)(nil)
