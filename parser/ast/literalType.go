package ast

type LiteralType struct {
	BaseNode

	structType  *StructType
	arrayType   *ArrayType
	elementType ElementType
	sliceType   *SliceType
	mapType     *MapType
	typeName    *TypeName
}

func (s *LiteralType) StructType() *StructType {
	return s.structType
}

func (s *LiteralType) SetStructType(structType *StructType) {
	s.structType = structType
}

func (s *LiteralType) ArrayType() *ArrayType {
	return s.arrayType
}

func (s *LiteralType) SetArrayType(arrayType *ArrayType) {
	s.arrayType = arrayType
}

func (s *LiteralType) ElementType() ElementType {
	return s.elementType
}

func (s *LiteralType) SetElementType(elementType ElementType) {
	s.elementType = elementType
}

func (s *LiteralType) SliceType() *SliceType {
	return s.sliceType
}

func (s *LiteralType) SetSliceType(sliceType *SliceType) {
	s.sliceType = sliceType
}

func (s *LiteralType) MapType() *MapType {
	return s.mapType
}

func (s *LiteralType) SetMapType(mapType *MapType) {
	s.mapType = mapType
}

func (s *LiteralType) TypeName() *TypeName {
	return s.typeName
}

func (s *LiteralType) SetTypeName(typeName *TypeName) {
	s.typeName = typeName
}

func (s *LiteralType) codeBuilder() *CodeBuilder {
	if s.structType != nil {
		return s.structType.codeBuilder()
	}
	if s.arrayType != nil {
		return s.arrayType.codeBuilder()
	}
	if s.elementType != nil {
		return NewCodeBuilder().appendString("[...]").appendNode(s.elementType)
	}
	if s.sliceType != nil {
		return s.sliceType.codeBuilder()
	}
	if s.mapType != nil {
		return s.mapType.codeBuilder()
	}
	return s.typeName.codeBuilder()
}

func (s *LiteralType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *LiteralType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*LiteralType)(nil)
