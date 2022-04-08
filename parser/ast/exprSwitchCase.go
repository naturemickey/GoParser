package ast

type ExprSwitchCase struct {
	BaseNode
	expressionList *ExpressionList
}

func NewExprSwitchCase(expressionList *ExpressionList) *ExprSwitchCase {
	return &ExprSwitchCase{expressionList: expressionList}
}

func (s *ExprSwitchCase) ExpressionList() *ExpressionList {
	return s.expressionList
}

func (s *ExprSwitchCase) SetExpressionList(expressionList *ExpressionList) {
	s.expressionList = expressionList
}

func (s *ExprSwitchCase) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.expressionList != nil {
		cb.appendString("case ").appendNode(s.expressionList)
	} else {
		cb.appendString("default")
	}
	return cb
}

func (s *ExprSwitchCase) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ExprSwitchCase) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ExprSwitchCase)(nil)
