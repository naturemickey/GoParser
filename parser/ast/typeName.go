package ast

type TypeName struct {
	BaseNode
}

func (s *TypeName) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeName) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeName) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeName)(nil)
