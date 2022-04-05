package ast

type Expression struct {
	BaseNode

	primaryExpr     *PrimaryExpr
	expression      *Expression
	expressionLeft  *Expression
	expressionRight *Expression

	unaryOp    string
	mulOp      string
	addOp      string
	relOp      string
	logicalAnd string
	logicalOr  string
}

func (s *Expression) _Element_() {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) _Key_() {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) PrimaryExpr() *PrimaryExpr {
	return s.primaryExpr
}

func (s *Expression) SetPrimaryExpr(primaryExpr *PrimaryExpr) {
	s.primaryExpr = primaryExpr
}

func (s *Expression) Expression() *Expression {
	return s.expression
}

func (s *Expression) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *Expression) ExpressionLeft() *Expression {
	return s.expressionLeft
}

func (s *Expression) SetExpressionLeft(expressionLeft *Expression) {
	s.expressionLeft = expressionLeft
}

func (s *Expression) ExpressionRight() *Expression {
	return s.expressionRight
}

func (s *Expression) SetExpressionRight(expressionRight *Expression) {
	s.expressionRight = expressionRight
}

func (s *Expression) UnaryOp() string {
	return s.unaryOp
}

func (s *Expression) SetUnaryOp(unaryOp string) {
	s.unaryOp = unaryOp
}

func (s *Expression) MulOp() string {
	return s.mulOp
}

func (s *Expression) SetMulOp(mulOp string) {
	s.mulOp = mulOp
}

func (s *Expression) AddOp() string {
	return s.addOp
}

func (s *Expression) SetAddOp(addOp string) {
	s.addOp = addOp
}

func (s *Expression) RelOp() string {
	return s.relOp
}

func (s *Expression) SetRelOp(relOp string) {
	s.relOp = relOp
}

func (s *Expression) LogicalAnd() string {
	return s.logicalAnd
}

func (s *Expression) SetLogicalAnd(logicalAnd string) {
	s.logicalAnd = logicalAnd
}

func (s *Expression) LogicalOr() string {
	return s.logicalOr
}

func (s *Expression) SetLogicalOr(logicalOr string) {
	s.logicalOr = logicalOr
}

func (s *Expression) _ArrayLength_() {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) _SimpleStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) _ExpressionStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) codeBuilder() *CodeBuilder {
	if s.primaryExpr != nil {
		return s.primaryExpr.codeBuilder()
	}
	cb := NewCodeBuilder()
	if s.unaryOp != "" {
		return cb.appendString(s.unaryOp).appendNode(s.expression)
	}

	if s.mulOp != "" {
		return cb.appendNode(s.expressionLeft).blank().appendString(s.mulOp).blank().appendNode(s.expressionRight)
	}

	if s.addOp != "" {
		return cb.appendNode(s.expressionLeft).blank().appendString(s.addOp).blank().appendNode(s.expressionRight)
	}

	if s.relOp != "" {
		return cb.appendNode(s.expressionLeft).blank().appendString(s.relOp).blank().appendNode(s.expressionRight)
	}

	if s.logicalAnd != "" {
		return cb.appendNode(s.expressionLeft).blank().appendString(s.logicalAnd).blank().appendNode(s.expressionRight)
	}
	return cb.appendNode(s.expressionLeft).blank().appendString(s.logicalOr).blank().appendNode(s.expressionRight)
}

func (s *Expression) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) String() string {
	return s.codeBuilder().String()
}

var _ ExpressionStmt = (*Expression)(nil)
var _ ArrayLength = (*Expression)(nil)
var _ Key = (*Expression)(nil)
var _ Element = (*Expression)(nil)
