package ast

type TypeName struct {
}

func (s TypeName) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeName) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeName)(nil)
