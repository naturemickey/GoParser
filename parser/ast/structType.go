package ast

type StructType struct {
}

func (s StructType) String() {
	//TODO implement me
	panic("implement me")
}

func (s StructType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*StructType)(nil)
