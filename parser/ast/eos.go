package ast

type Eos struct {
}

func (s Eos) String() {
	//TODO implement me
	panic("implement me")
}

func (s Eos) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Eos)(nil)
