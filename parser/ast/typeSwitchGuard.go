package ast

type TypeSwitchGuard struct {
}

func (s TypeSwitchGuard) String() {
	//TODO implement me
	panic("implement me")
}

func (s TypeSwitchGuard) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*TypeSwitchGuard)(nil)
