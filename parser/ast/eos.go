package ast

type Eos struct {
	BaseNode

	literal string
}

func NewEos(literal string) *Eos {
	return &Eos{literal: literal}
}

func (s *Eos) Literal() string {
	return s.literal
}

func (s *Eos) SetLiteral(literal string) {
	s.literal = literal
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
