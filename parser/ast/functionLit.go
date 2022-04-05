package ast

type FunctionLit struct {
}

func (s FunctionLit) String() {
	//TODO implement me
	panic("implement me")
}

func (s FunctionLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*FunctionLit)(nil)
