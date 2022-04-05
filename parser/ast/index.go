package ast

type Index struct {
}

func (s Index) String() {
	//TODO implement me
	panic("implement me")
}

func (s Index) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Index)(nil)
