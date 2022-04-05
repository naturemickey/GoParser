package ast

type PointerType struct {
	BaseNode
}

func (s *PointerType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *PointerType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *PointerType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*PointerType)(nil)
