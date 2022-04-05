package ast

type VarSpec struct {
	BaseNode
	identifierList *IdentifierList
	type_          *Type_
	expressionList *ExpressionList
}

func (s *VarSpec) IdentifierList() *IdentifierList {
	return s.identifierList
}

func (s *VarSpec) SetIdentifierList(identifierList *IdentifierList) {
	s.identifierList = identifierList
}

func (s *VarSpec) Type_() *Type_ {
	return s.type_
}

func (s *VarSpec) SetType_(type_ *Type_) {
	s.type_ = type_
}

func (s *VarSpec) ExpressionList() *ExpressionList {
	return s.expressionList
}

func (s *VarSpec) SetExpressionList(expressionList *ExpressionList) {
	s.expressionList = expressionList
}

func (s *VarSpec) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendNode(s.identifierList).blank()
	if s.type_ != nil {
		cb.appendNode(s.type_).blank()
	}
	if s.expressionList != nil {
		cb.appendString("= ").appendNode(s.expressionList)
	}
	return cb
}

func (s *VarSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *VarSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*VarSpec)(nil)
