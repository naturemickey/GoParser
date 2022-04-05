package ast

type Slice struct {
}

func (s Slice) String() {
	//TODO implement me
	panic("implement me")
}

func (s Slice) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Slice)(nil)
