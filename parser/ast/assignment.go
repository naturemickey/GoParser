package ast

type Assignment struct {
}

func (s Assignment) String() {
	//TODO implement me
	panic("implement me")
}

func (s Assignment) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Assignment)(nil)
