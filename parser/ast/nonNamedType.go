package ast

type NonNamedType struct {
}

func (s NonNamedType) String() {
	//TODO implement me
	panic("implement me")
}

func (s NonNamedType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*NonNamedType)(nil)
