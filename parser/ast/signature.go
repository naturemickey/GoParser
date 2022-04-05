package ast

type Signature struct {
}

func (s Signature) String() {
	//TODO implement me
	panic("implement me")
}

func (s Signature) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Signature)(nil)
