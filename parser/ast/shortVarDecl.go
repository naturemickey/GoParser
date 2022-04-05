package ast

type ShortValDecl struct {
}

func (s ShortValDecl) String() {
	//TODO implement me
	panic("implement me")
}

func (s ShortValDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ShortValDecl)(nil)
