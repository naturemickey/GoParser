package ast

type FunctionDecl struct {
}

func (s FunctionDecl) String() {
	//TODO implement me
	panic("implement me")
}

func (s FunctionDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*FunctionDecl)(nil)
