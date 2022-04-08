package ast

type ConstSpec struct {
	BaseNode
	identifierList *IdentifierList
	type_          *Type_
	expressionList *ExpressionList
}

func NewConstSpec(identifierList *IdentifierList, type_ *Type_, expressionList *ExpressionList) *ConstSpec {
	return &ConstSpec{identifierList: identifierList, type_: type_, expressionList: expressionList}
}

func (s *ConstSpec) IdentifierList() *IdentifierList {
	return s.identifierList
}

func (s *ConstSpec) SetIdentifierList(identifierList *IdentifierList) {
	s.identifierList = identifierList
}

func (s *ConstSpec) Type_() *Type_ {
	return s.type_
}

func (s *ConstSpec) SetType_(type_ *Type_) {
	s.type_ = type_
}

func (s *ConstSpec) ExpressionList() *ExpressionList {
	return s.expressionList
}

func (s *ConstSpec) SetExpressionList(expressionList *ExpressionList) {
	s.expressionList = expressionList
}

func (s *ConstSpec) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendNode(s.identifierList).blank()

	if s.expressionList != nil {
		if s.type_ != nil {
			cb.appendNode(s.type_).blank()
		}
		cb.appendString("= ")
		cb.appendNode(s.expressionList)
	}
	return cb
}

func (s *ConstSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ConstSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ConstSpec)(nil)
