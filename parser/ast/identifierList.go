package ast

type IdentifierList struct {
}

func (s IdentifierList) String() {
	//TODO implement me
	panic("implement me")
}

func (s IdentifierList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*IdentifierList)(nil)
