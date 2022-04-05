package ast

type Parameters struct {
	BaseNode
}

func (s *Parameters) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Parameters) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Parameters) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Parameters)(nil)
