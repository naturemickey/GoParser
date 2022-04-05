package ast

type ReceiverType struct {
}

func (s ReceiverType) String() {
	//TODO implement me
	panic("implement me")
}

func (s ReceiverType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ReceiverType)(nil)
