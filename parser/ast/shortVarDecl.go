package ast

type ShortValDecl struct {
	BaseNode
	identifierList *IdentifierList
	expressionList *ExpressionList
}

func NewShortValDecl(identifierList *IdentifierList, expressionList *ExpressionList) *ShortValDecl {
	return &ShortValDecl{identifierList: identifierList, expressionList: expressionList}
}

func (s *ShortValDecl) IdentifierList() *IdentifierList {
	return s.identifierList
}

func (s *ShortValDecl) SetIdentifierList(identifierList *IdentifierList) {
	s.identifierList = identifierList
}

func (s *ShortValDecl) ExpressionList() *ExpressionList {
	return s.expressionList
}

func (s *ShortValDecl) SetExpressionList(expressionList *ExpressionList) {
	s.expressionList = expressionList
}

func (s *ShortValDecl) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *ShortValDecl) _SimpleStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *ShortValDecl) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendNode(s.identifierList).appendString(" := ").appendNode(s.expressionList)
}

func (s *ShortValDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ShortValDecl) String() string {
	return s.codeBuilder().String()
}

var _ SimpleStmt = (*ShortValDecl)(nil)
