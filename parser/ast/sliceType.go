package ast

type SliceType struct {
	BaseNode

	elementType ElementType
}

func NewSliceType(elementType ElementType) *SliceType {
	return &SliceType{elementType: elementType}
}

func (s *SliceType) ElementType() ElementType {
	return s.elementType
}

func (s *SliceType) SetElementType(elementType ElementType) {
	s.elementType = elementType
}

func (s *SliceType) _TypeLit_() {
	//TODO implement me
	panic("implement me")
}

func (s *SliceType) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("[]").appendNode(s.elementType)
}

func (s *SliceType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SliceType) String() string {
	return s.codeBuilder().String()
}

var _ TypeLit = (*SliceType)(nil)
