package ast

type KeyedElement struct {
	BaseNode
}

func (s *KeyedElement) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *KeyedElement) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *KeyedElement) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*KeyedElement)(nil)
