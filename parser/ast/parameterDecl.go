package ast

type ParameterDecl struct {
}

func (s ParameterDecl) String() {
	//TODO implement me
	panic("implement me")
}

func (s ParameterDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ParameterDecl)(nil)
