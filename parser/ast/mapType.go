package ast

type MapType struct {
	BaseNode

	type_       *Type_
	elementType ElementType
}

func (s *MapType) Type_() *Type_ {
	return s.type_
}

func (s *MapType) SetType_(type_ *Type_) {
	s.type_ = type_
}

func (s *MapType) ElementType() ElementType {
	return s.elementType
}

func (s *MapType) SetElementType(elementType ElementType) {
	s.elementType = elementType
}

func (s *MapType) _TypeLit_() {
	//TODO implement me
	panic("implement me")
}

func (s *MapType) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("map[").appendNode(s.type_).appendString("]").appendNode(s.elementType)
}

func (s *MapType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *MapType) String() string {
	return s.codeBuilder().String()
}

var _ TypeLit = (*MapType)(nil)
