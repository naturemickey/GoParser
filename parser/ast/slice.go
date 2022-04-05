package ast

type Slice struct {
	BaseNode
}

func (s *Slice) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Slice) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Slice) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Slice)(nil)
