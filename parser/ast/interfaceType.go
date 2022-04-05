package ast

type InterfaceType struct {
}

func (s InterfaceType) String() {
	//TODO implement me
	panic("implement me")
}

func (s InterfaceType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*InterfaceType)(nil)
