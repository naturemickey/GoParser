package ast

type OperandName struct {
}

func (s OperandName) String() {
	//TODO implement me
	panic("implement me")
}

func (s OperandName) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*OperandName)(nil)
