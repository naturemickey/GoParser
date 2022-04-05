package ast

type Result struct {
}

func (s Result) String() {
	//TODO implement me
	panic("implement me")
}

func (s Result) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Result)(nil)
