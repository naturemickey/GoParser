package ast

type Type_ struct {
}

func (s Type_) String() {
	//TODO implement me
	panic("implement me")
}

func (s Type_) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Type_)(nil)
