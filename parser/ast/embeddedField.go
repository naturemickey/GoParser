package ast

type EmbeddedField struct {
}

func (s EmbeddedField) String() {
	//TODO implement me
	panic("implement me")
}

func (s EmbeddedField) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*EmbeddedField)(nil)
