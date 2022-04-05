package ast

type CompositeLit struct {
	BaseNode

	literalType  *LiteralType
	literalValue *LiteralValue
}

func (s *CompositeLit) _Literal_() {
	//TODO implement me
	panic("implement me")
}

func (s *CompositeLit) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendNode(s.literalType).appendNode(s.literalValue)
}

func (s *CompositeLit) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *CompositeLit) String() string {
	return s.codeBuilder().String()
}

var _ Literal = (*CompositeLit)(nil)
