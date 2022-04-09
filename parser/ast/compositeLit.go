package ast

type CompositeLit struct {
	BaseNode

	literalType  *LiteralType
	literalValue *LiteralValue
}

func NewCompositeLit(literalType *LiteralType, literalValue *LiteralValue) *CompositeLit {
	return &CompositeLit{literalType: literalType, literalValue: literalValue}
}

func (s *CompositeLit) LiteralType() *LiteralType {
	return s.literalType
}

func (s *CompositeLit) SetLiteralType(literalType *LiteralType) {
	s.literalType = literalType
}

func (s *CompositeLit) LiteralValue() *LiteralValue {
	return s.literalValue
}

func (s *CompositeLit) SetLiteralValue(literalValue *LiteralValue) {
	s.literalValue = literalValue
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
