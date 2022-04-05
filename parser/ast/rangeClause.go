package ast

type RangeClause struct {
	BaseNode

	expressionList *ExpressionList
	identifierList *IdentifierList
	expression     *Expression
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
