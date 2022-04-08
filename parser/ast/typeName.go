package ast

type TypeName struct {
	BaseNode
	qualifiedIdent *QualifiedIdent
	id             string
}

func NewTypeName(qualifiedIdent *QualifiedIdent, id string) *TypeName {
	return &TypeName{qualifiedIdent: qualifiedIdent, id: id}
}

func (s *TypeName) QualifiedIdent() *QualifiedIdent {
	return s.qualifiedIdent
}

func (s *TypeName) SetQualifiedIdent(qualifiedIdent *QualifiedIdent) {
	s.qualifiedIdent = qualifiedIdent
}

func (s *TypeName) Id() string {
	return s.id
}

func (s *TypeName) SetId(id string) {
	s.id = id
}

func (s *TypeName) _IMethodSpecOrTypeName_() {
	//TODO implement me
	panic("implement me")
}

func (s *TypeName) codeBuilder() *CodeBuilder {
	if s.qualifiedIdent != nil {
		return s.qualifiedIdent.codeBuilder()
	}
	return NewCodeBuilder().appendString(s.id)
}

func (s *TypeName) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeName) String() string {
	return s.codeBuilder().String()
}

var _ IMethodSpecOrTypeName = (*TypeName)(nil)
