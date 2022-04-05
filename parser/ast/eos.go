package ast

type Eos struct {
	BaseNode
}

func (s *Eos) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Eos) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Eos) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Eos)(nil)
