package ast

type ImportPath struct {
}

func (s ImportPath) String() {
	//TODO implement me
	panic("implement me")
}

func (s ImportPath) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ImportPath)(nil)
