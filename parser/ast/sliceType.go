package ast

type SliceType struct {
	BaseNode
}

func (s *SliceType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *SliceType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SliceType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*SliceType)(nil)
