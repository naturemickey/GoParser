package ast

type QualifiedIdent struct {
	BaseNode
}

func (s *QualifiedIdent) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *QualifiedIdent) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *QualifiedIdent) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*QualifiedIdent)(nil)
