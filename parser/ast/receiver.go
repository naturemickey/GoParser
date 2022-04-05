package ast

type Receiver struct {
}

func (s Receiver) String() {
	//TODO implement me
	panic("implement me")
}

func (s Receiver) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Receiver)(nil)
