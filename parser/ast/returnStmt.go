package ast

type ReturnStmt struct {
	BaseNode
	expressionList *ExpressionList
}

func NewReturnStmt(expressionList *ExpressionList) *ReturnStmt {
	return &ReturnStmt{expressionList: expressionList}
}

func (s *ReturnStmt) ExpressionList() *ExpressionList {
	return s.expressionList
}

func (s *ReturnStmt) SetExpressionList(expressionList *ExpressionList) {
	s.expressionList = expressionList
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
