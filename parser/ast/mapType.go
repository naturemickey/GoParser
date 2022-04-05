package ast

type MapType struct {
}

func (s MapType) String() {
	//TODO implement me
	panic("implement me")
}

func (s MapType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*MapType)(nil)
