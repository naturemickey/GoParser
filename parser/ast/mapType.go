package ast

type MapType struct {
	BaseNode
}

func (s *MapType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *MapType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *MapType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*MapType)(nil)
