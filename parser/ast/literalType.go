package ast

type LiteralType struct {
	BaseNode
}

func (s *LiteralType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *LiteralType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *LiteralType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*LiteralType)(nil)
