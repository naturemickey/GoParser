package ast

type Literal struct {
}

func (s Literal) String() {
	//TODO implement me
	panic("implement me")
}

func (s Literal) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Literal)(nil)
