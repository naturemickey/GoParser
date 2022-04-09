package ast

type Expression struct {
	BaseNode

	primaryExpr *PrimaryExpr
	expression  *Expression
	expression2 *Expression

	unaryOp    string
	mulOp      string
	addOp      string
	relOp      string
	logicalAnd string
	logicalOr  string
}

type ExpressionBuilder struct {
	primaryExpr *PrimaryExpr
	expression  *Expression
	expression2 *Expression

	unaryOp    string
	mulOp      string
	addOp      string
	relOp      string
	logicalAnd string
	logicalOr  string
}

func NewExpressionBuilder() *ExpressionBuilder {
	return &ExpressionBuilder{}
}

func (b *ExpressionBuilder) SetPrimaryExpr(primaryExpr *PrimaryExpr) *ExpressionBuilder {
	b.primaryExpr = primaryExpr
	return b
}

func (b *ExpressionBuilder) SetExpression(expression *Expression) *ExpressionBuilder {
	b.expression = expression
	return b
}

func (b *ExpressionBuilder) SetExpression2(expression2 *Expression) *ExpressionBuilder {
	b.expression2 = expression2
	return b
}

func (b *ExpressionBuilder) SetUnaryOp(unaryOp string) *ExpressionBuilder {
	b.unaryOp = unaryOp
	return b
}

func (b *ExpressionBuilder) SetMulOp(mulOp string) *ExpressionBuilder {
	b.mulOp = mulOp
	return b
}

func (b *ExpressionBuilder) SetAddOp(addOp string) *ExpressionBuilder {
	b.addOp = addOp
	return b
}

func (b *ExpressionBuilder) SetRelOp(relOp string) *ExpressionBuilder {
	b.relOp = relOp
	return b
}

func (b *ExpressionBuilder) SetLogicalAnd(logicalAnd string) *ExpressionBuilder {
	b.logicalAnd = logicalAnd
	return b
}

func (b *ExpressionBuilder) SetLogicalOr(logicalOr string) *ExpressionBuilder {
	b.logicalOr = logicalOr
	return b
}

func (b *ExpressionBuilder) Build() *Expression {
	return NewExpression(b.primaryExpr, b.expression, b.expression2, b.unaryOp, b.mulOp, b.addOp, b.relOp, b.logicalAnd, b.logicalOr)
}

func NewExpression(primaryExpr *PrimaryExpr, expression *Expression, expression2 *Expression, unaryOp string, mulOp string, addOp string, relOp string, logicalAnd string, logicalOr string) *Expression {
	return &Expression{primaryExpr: primaryExpr, expression: expression, expression2: expression2, unaryOp: unaryOp, mulOp: mulOp, addOp: addOp, relOp: relOp, logicalAnd: logicalAnd, logicalOr: logicalOr}
}

func (s *Expression) Expression2() *Expression {
	return s.expression2
}

func (s *Expression) SetExpression2(expression2 *Expression) {
	s.expression2 = expression2
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
		return cb.appendNode(s.expression).blank().appendString(s.mulOp).blank().appendNode(s.expression2)
	}
	if s.addOp != "" {
		return cb.appendNode(s.expression).blank().appendString(s.addOp).blank().appendNode(s.expression2)
	}
	if s.relOp != "" {
		return cb.appendNode(s.expression).blank().appendString(s.relOp).blank().appendNode(s.expression2)
	}
	if s.logicalAnd != "" {
		return cb.appendNode(s.expression).blank().appendString(s.logicalAnd).blank().appendNode(s.expression2)
	}
	return cb.appendNode(s.expression).blank().appendString(s.logicalOr).blank().appendNode(s.expression2)
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
