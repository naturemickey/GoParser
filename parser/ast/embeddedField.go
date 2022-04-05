package ast

type EmbeddedField struct {
	BaseNode
}

func (s *EmbeddedField) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *EmbeddedField) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *EmbeddedField) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*EmbeddedField)(nil)
