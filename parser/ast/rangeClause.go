package ast

type RangeClause struct {
	BaseNode

	expressionList *ExpressionList
	identifierList *IdentifierList
	expression     *Expression
}

func NewRangeClause(expressionList *ExpressionList, identifierList *IdentifierList, expression *Expression) *RangeClause {
	return &RangeClause{expressionList: expressionList, identifierList: identifierList, expression: expression}
}

func (s *RangeClause) ExpressionList() *ExpressionList {
	return s.expressionList
}

func (s *RangeClause) SetExpressionList(expressionList *ExpressionList) {
	s.expressionList = expressionList
}

func (s *RangeClause) IdentifierList() *IdentifierList {
	return s.identifierList
}

func (s *RangeClause) SetIdentifierList(identifierList *IdentifierList) {
	s.identifierList = identifierList
}

func (s *RangeClause) Expression() *Expression {
	return s.expression
}

func (s *RangeClause) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *RangeClause) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.expressionList != nil {
		cb.appendNode(s.expressionList).appendString(" = ")
	} else {
		cb.appendNode(s.identifierList).appendString(" := ")
	}
	cb.appendString("range ").appendNode(s.expression)
	return cb
}

func (s *RangeClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *RangeClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*RangeClause)(nil)
