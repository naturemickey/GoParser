package ast

type ArrayType struct {
	BaseNode

	arrayLength ArrayLength
	elementType ElementType
}

func NewArrayType(arrayLength ArrayLength, elementType ElementType) *ArrayType {
	return &ArrayType{arrayLength: arrayLength, elementType: elementType}
}

func (s *ArrayType) ArrayLength() ArrayLength {
	return s.arrayLength
}

func (s *ArrayType) SetArrayLength(arrayLength ArrayLength) {
	s.arrayLength = arrayLength
}

func (s *ArrayType) ElementType() ElementType {
	return s.elementType
}

func (s *ArrayType) SetElementType(elementType ElementType) {
	s.elementType = elementType
}

func (s *ArrayType) _TypeLit_() {
	//TODO implement me
	panic("implement me")
}

func (s *ArrayType) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("[").appendNode(s.arrayLength).appendString("]").appendNode(s.elementType)
}

func (s *ArrayType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ArrayType) String() string {
	return s.codeBuilder().String()
}

var _ TypeLit = (*ArrayType)(nil)
