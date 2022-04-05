package ast

type CommCase struct {
}

func (s CommCase) String() {
	//TODO implement me
	panic("implement me")
}

func (s CommCase) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*CommCase)(nil)
