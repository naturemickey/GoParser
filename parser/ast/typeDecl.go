package ast

type TypeDecl struct {
}

func (s TypeDecl) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeDecl)(nil)
