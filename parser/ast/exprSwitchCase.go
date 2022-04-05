package ast

type ExprSwitchCase struct {
	BaseNode
	expressionList *ExpressionList
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
