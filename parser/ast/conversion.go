package ast

type Conversion struct {
}

func (s Conversion) String() {
	//TODO implement me
	panic("implement me")
}

func (s Conversion) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Conversion)(nil)
