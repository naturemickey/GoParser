package ast

type Declaration struct {
}

func (s Declaration) String() {
	//TODO implement me
	panic("implement me")
}

func (s Declaration) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Declaration)(nil)
