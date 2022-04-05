package ast

type FieldDecl struct {
}

func (s FieldDecl) String() {
	//TODO implement me
	panic("implement me")
}

func (s FieldDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*FieldDecl)(nil)
