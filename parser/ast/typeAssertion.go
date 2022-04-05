package ast

type TypeAssertion struct {
}

func (s TypeAssertion) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeAssertion) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeAssertion)(nil)
