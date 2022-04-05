package ast

type ArrayType struct {
}

func (s ArrayType) String() {
	//TODO implement me
	panic("implement me")
}

func (s ArrayType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ArrayType)(nil)
