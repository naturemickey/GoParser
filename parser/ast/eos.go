package ast

type Eos struct {
	BaseNode

	literal string
}

func (s *Eos) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString(s.literal)
}

func (s *Eos) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Eos) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Eos)(nil)
