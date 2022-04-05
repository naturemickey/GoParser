package ast

type ReturnStmt struct {
	BaseNode
	expressionList *ExpressionList
}

func (s *ReturnStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *ReturnStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("return")
	if s.expressionList != nil {
		cb.blank().appendNode(s.expressionList)
	}
	return cb
}

func (s *ReturnStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ReturnStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*ReturnStmt)(nil)
