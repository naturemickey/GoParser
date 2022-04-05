package ast

type MethodDecl struct {
}

func (s MethodDecl) String() {
	//TODO implement me
	panic("implement me")
}

func (s MethodDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*MethodDecl)(nil)
