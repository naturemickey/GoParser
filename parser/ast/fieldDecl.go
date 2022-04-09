package ast

type FieldDecl struct {
	BaseNode

	identifierList *IdentifierList
	type_          *Type_
	embeddedField  *EmbeddedField

	tag string
}

func NewFieldDecl(identifierList *IdentifierList, type_ *Type_, embeddedField *EmbeddedField, tag string) *FieldDecl {
	return &FieldDecl{identifierList: identifierList, type_: type_, embeddedField: embeddedField, tag: tag}
}

func (s *FieldDecl) IdentifierList() *IdentifierList {
	return s.identifierList
}

func (s *FieldDecl) SetIdentifierList(identifierList *IdentifierList) {
	s.identifierList = identifierList
}

func (s *FieldDecl) Type_() *Type_ {
	return s.type_
}

func (s *FieldDecl) SetType_(type_ *Type_) {
	s.type_ = type_
}

func (s *FieldDecl) EmbeddedField() *EmbeddedField {
	return s.embeddedField
}

func (s *FieldDecl) SetEmbeddedField(embeddedField *EmbeddedField) {
	s.embeddedField = embeddedField
}

func (s *FieldDecl) Tag() string {
	return s.tag
}

func (s *FieldDecl) SetTag(tag string) {
	s.tag = tag
}

func (s *FieldDecl) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.embeddedField != nil {
		cb.appendNode(s.embeddedField)
	} else {
		cb.appendNode(s.identifierList).blank().appendNode(s.type_)
	}
	if s.tag != "" {
		cb.blank().appendString(s.tag)
	}
	return cb
}

func (s *FieldDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FieldDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*FieldDecl)(nil)
