package ast

type Block struct {
}

func (s Block) String() {
	//TODO implement me
	panic("implement me")
}

func (s Block) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Block)(nil)
