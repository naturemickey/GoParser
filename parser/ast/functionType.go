package ast

type FunctionType struct {
}

func (s FunctionType) String() {
	//TODO implement me
	panic("implement me")
}

func (s FunctionType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*FunctionType)(nil)
