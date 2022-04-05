package ast

type Signature struct {
}

func (s *Signature) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Signature) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Signature) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Signature)(nil)
