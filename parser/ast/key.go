package ast

type Key struct {
	BaseNode
}

func (s *Key) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Key) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Key) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Key)(nil)
