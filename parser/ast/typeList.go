package ast

type TypeList struct {
}

func (s TypeList) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeList)(nil)
