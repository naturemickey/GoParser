package ast

type VarSpec struct {
}

func (s VarSpec) String() {
	//TODO implement me
	panic("implement me")
}

func (s VarSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*VarSpec)(nil)
