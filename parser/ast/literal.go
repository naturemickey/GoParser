package ast

type Literal struct {
	BaseNode
}

func (s *Literal) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Literal) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Literal) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Literal)(nil)
