package ast

type EmbeddedField struct {
	BaseNode

	star     bool
	typeName *TypeName
}

func NewEmbeddedField(star bool, typeName *TypeName) *EmbeddedField {
	return &EmbeddedField{star: star, typeName: typeName}
}

func (s *EmbeddedField) Star() bool {
	return s.star
}

func (s *EmbeddedField) SetStar(star bool) {
	s.star = star
}

func (s *EmbeddedField) TypeName() *TypeName {
	return s.typeName
}

func (s *EmbeddedField) SetTypeName(typeName *TypeName) {
	s.typeName = typeName
}

func (s *EmbeddedField) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.star {
		cb.appendString("*")
	}
	cb.appendNode(s.typeName)
	return cb
}

func (s *EmbeddedField) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *EmbeddedField) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*EmbeddedField)(nil)
