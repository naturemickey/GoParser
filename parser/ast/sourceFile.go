package ast

type SourceFile struct {
}

func (s SourceFile) String() {
	//TODO implement me
	panic("implement me")
}

func (s SourceFile) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*SourceFile)(nil)
