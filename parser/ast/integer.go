package ast

type Integer struct {
	BaseNode

	intType IntegerLitType
	literal string
}

func NewInteger(intType IntegerLitType, literal string) *Integer {
	return &Integer{intType: intType, literal: literal}
}

func (s *Integer) IntType() IntegerLitType {
	return s.intType
}

func (s *Integer) SetIntType(intType IntegerLitType) {
	s.intType = intType
}

func (s *Integer) Literal() string {
	return s.literal
}

func (s *Integer) SetLiteral(literal string) {
	s.literal = literal
}

type IntegerLitType int

const (
	DECIMAL_LIT = iota
	BINARY_LIT
	OCTAL_LIT
	HEX_LIT
	IMAGINARY_LIT
	RUNE_LIT
)

func (s *Integer) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString(s.literal)
}

func (s *Integer) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Integer) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Integer)(nil)
