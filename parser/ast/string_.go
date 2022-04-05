package ast

type String_ struct {
}

func (s String_) String() {
	//TODO implement me
	panic("implement me")
}

func (s String_) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*String_)(nil)
