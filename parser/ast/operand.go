package ast

type Operand struct {
}

func (s Operand) String() {
	//TODO implement me
	panic("implement me")
}

func (s Operand) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Operand)(nil)
