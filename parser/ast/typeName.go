package ast

type TypeName struct {
	BaseNode
	qualifiedIdent *QualifiedIdent
	id             string
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
