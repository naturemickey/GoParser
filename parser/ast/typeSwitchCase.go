package ast

type TypeSwitchCase struct {
}

func (s TypeSwitchCase) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeSwitchCase) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeSwitchCase)(nil)
