package ast

type Type_ struct {
	BaseNode

	typeName *TypeName
	typeLit  TypeLit
	type_    *Type_
}

func NewType_(typeName *TypeName, typeLit TypeLit, type_ *Type_) *Type_ {
	return &Type_{typeName: typeName, typeLit: typeLit, type_: type_}
}

func (s *Type_) TypeName() *TypeName {
	return s.typeName
}

func (s *Type_) SetTypeName(typeName *TypeName) {
	s.typeName = typeName
}

func (s *Type_) TypeLit() TypeLit {
	return s.typeLit
}

func (s *Type_) SetTypeLit(typeLit TypeLit) {
	s.typeLit = typeLit
}

func (s *Type_) Type_() *Type_ {
	return s.type_
}

func (s *Type_) SetType_(type_ *Type_) {
	s.type_ = type_
}

func (s *Type_) _ReceiverType_() {
	//TODO implement me
	panic("implement me")
}

func (s *Type_) _Result_() {
	//TODO implement me
	panic("implement me")
}

func (s *Type_) _ElementType_() {
	//TODO implement me
	panic("implement me")
}

func (s *Type_) codeBuilder() *CodeBuilder {
	if s.typeName != nil {
		return s.typeName.codeBuilder()
	}
	if s.typeLit != nil {
		return s.typeLit.codeBuilder()
	}
	return NewCodeBuilder().appendString("(").appendNode(s.type_).appendString(")")
}

func (s *Type_) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Type_) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Type_)(nil)
var _ ElementType = (*Type_)(nil)
var _ Result = (*Type_)(nil)
var _ ReceiverType = (*Type_)(nil)
