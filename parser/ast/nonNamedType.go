package ast

type NonNamedType struct {
	BaseNode

	typeLit      TypeLit
	nonNamedType *NonNamedType
}

func (s *NonNamedType) TypeLit() TypeLit {
	return s.typeLit
}

func (s *NonNamedType) SetTypeLit(typeLit TypeLit) {
	s.typeLit = typeLit
}

func (s *NonNamedType) NonNamedType() *NonNamedType {
	return s.nonNamedType
}

func (s *NonNamedType) SetNonNamedType(nonNamedType *NonNamedType) {
	s.nonNamedType = nonNamedType
}

func (s *NonNamedType) codeBuilder() *CodeBuilder {
	if s.typeLit != nil {
		return s.typeLit.codeBuilder()
	}
	return NewCodeBuilder().appendString("(").appendNode(s.nonNamedType).appendString(")")
}

func (s *NonNamedType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *NonNamedType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*NonNamedType)(nil)
