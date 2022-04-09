package ast

type ParameterDecl struct {
	BaseNode

	identifierList *IdentifierList
	ellipsis       bool
	type_          *Type_
}

func NewParameterDecl(identifierList *IdentifierList, ellipsis bool, type_ *Type_) *ParameterDecl {
	return &ParameterDecl{identifierList: identifierList, ellipsis: ellipsis, type_: type_}
}

func (s *ParameterDecl) IdentifierList() *IdentifierList {
	return s.identifierList
}

func (s *ParameterDecl) SetIdentifierList(identifierList *IdentifierList) {
	s.identifierList = identifierList
}

func (s *ParameterDecl) Ellipsis() bool {
	return s.ellipsis
}

func (s *ParameterDecl) SetEllipsis(ellipsis bool) {
	s.ellipsis = ellipsis
}

func (s *ParameterDecl) Type_() *Type_ {
	return s.type_
}

func (s *ParameterDecl) SetType_(type_ *Type_) {
	s.type_ = type_
}

func (s *ParameterDecl) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendNode(s.identifierList).blank()
	if s.ellipsis {
		cb.appendString("... ")
	}
	cb.appendNode(s.type_)
	return cb
}

func (s *ParameterDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ParameterDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ParameterDecl)(nil)
