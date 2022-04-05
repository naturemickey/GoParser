package ast

type Receiver struct {
	BaseNode
}

func (s *Receiver) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Receiver) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Receiver) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Receiver)(nil)
