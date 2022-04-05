package ast

type LiteralValue struct {
	BaseNode
}

func (s *LiteralValue) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *LiteralValue) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *LiteralValue) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*LiteralValue)(nil)
