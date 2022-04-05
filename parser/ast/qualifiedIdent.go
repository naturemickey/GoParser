package ast

type QualifiedIdent struct {
}

func (s QualifiedIdent) String() {
	//TODO implement me
	panic("implement me")
}

func (s QualifiedIdent) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*QualifiedIdent)(nil)
