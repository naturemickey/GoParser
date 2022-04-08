package ast

type TypeSpec struct {
	BaseNode
	id    string
	type_ *Type_
}

func NewTypeSpec(id string, type_ *Type_) *TypeSpec {
	return &TypeSpec{id: id, type_: type_}
}

func (s *TypeSpec) Id() string {
	return s.id
}

func (s *TypeSpec) SetId(id string) {
	s.id = id
}

func (s *TypeSpec) Type_() *Type_ {
	return s.type_
}

func (s *TypeSpec) SetType_(type_ *Type_) {
	s.type_ = type_
}

func (s *TypeSpec) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString(s.id).blank().appendNode(s.type_)
}

func (s *TypeSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeSpec)(nil)
