package ast

type Arguments struct {
	BaseNode

	expressionList *ExpressionList
	nonNamedType   *NonNamedType
	ellipsis       bool
}

func NewArguments(expressionList *ExpressionList, nonNamedType *NonNamedType, ellipsis bool) *Arguments {
	return &Arguments{expressionList: expressionList, nonNamedType: nonNamedType, ellipsis: ellipsis}
}

func (s *Arguments) ExpressionList() *ExpressionList {
	return s.expressionList
}

func (s *Arguments) SetExpressionList(expressionList *ExpressionList) {
	s.expressionList = expressionList
}

func (s *Arguments) NonNamedType() *NonNamedType {
	return s.nonNamedType
}

func (s *Arguments) SetNonNamedType(nonNamedType *NonNamedType) {
	s.nonNamedType = nonNamedType
}

func (s *Arguments) Ellipsis() bool {
	return s.ellipsis
}

func (s *Arguments) SetEllipsis(ellipsis bool) {
	s.ellipsis = ellipsis
}

func (s *Arguments) String() string {
	return s.codeBuilder().String()
}

func (s *Arguments) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("(")
	if s.nonNamedType != nil {
		cb.appendNode(s.nonNamedType)
		if s.expressionList != nil {
			cb.appendString(", ").appendNode(s.expressionList)
		}
	} else if s.expressionList != nil {
		cb.appendNode(s.expressionList)
	}
	if s.ellipsis {
		cb.appendString("...")
	}
	cb.appendString(")")
	return cb
}

func (s *Arguments) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*Arguments)(nil)
