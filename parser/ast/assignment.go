package ast

type Assignment struct {
	BaseNode
	expressionListLeft  *ExpressionList
	assignOp            *AssignOp
	expressionListRight *ExpressionList
}

func (s *Assignment) ExpressionListLeft() *ExpressionList {
	return s.expressionListLeft
}

func (s *Assignment) SetExpressionListLeft(expressionListLeft *ExpressionList) {
	s.expressionListLeft = expressionListLeft
}

func (s *Assignment) AssignOp() *AssignOp {
	return s.assignOp
}

func (s *Assignment) SetAssignOp(assignOp *AssignOp) {
	s.assignOp = assignOp
}

func (s *Assignment) ExpressionListRight() *ExpressionList {
	return s.expressionListRight
}

func (s *Assignment) SetExpressionListRight(expressionListRight *ExpressionList) {
	s.expressionListRight = expressionListRight
}

func (s *Assignment) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *Assignment) _SimpleStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *Assignment) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendNode(s.expressionListLeft).blank().appendNode(s.assignOp).blank().appendNode(s.expressionListRight)
}

func (s *Assignment) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Assignment) String() string {
	return s.codeBuilder().String()
}

var _ SimpleStmt = (*Assignment)(nil)
