package ast

type Operand struct {
	BaseNode
}

func (s *Operand) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Operand) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Operand) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Operand)(nil)
