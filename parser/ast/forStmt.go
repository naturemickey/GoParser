package ast

type ForStmt struct {
	BaseNode

	expression  *Expression
	forClause   *ForClause
	rangeClause *RangeClause
	block       *Block
}

func NewForStmt(expression *Expression, forClause *ForClause, rangeClause *RangeClause, block *Block) *ForStmt {
	return &ForStmt{expression: expression, forClause: forClause, rangeClause: rangeClause, block: block}
}

func (s *ForStmt) Expression() *Expression {
	return s.expression
}

func (s *ForStmt) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *ForStmt) ForClause() *ForClause {
	return s.forClause
}

func (s *ForStmt) SetForClause(forClause *ForClause) {
	s.forClause = forClause
}

func (s *ForStmt) RangeClause() *RangeClause {
	return s.rangeClause
}

func (s *ForStmt) SetRangeClause(rangeClause *RangeClause) {
	s.rangeClause = rangeClause
}

func (s *ForStmt) Block() *Block {
	return s.block
}

func (s *ForStmt) SetBlock(block *Block) {
	s.block = block
}

func (s *ForStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *ForStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("for ")
	cb.appendNode(s.expression)
	cb.appendNode(s.forClause)
	cb.appendNode(s.rangeClause)
	cb.appendNode(s.block)
	return cb
}

func (s *ForStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ForStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*ForStmt)(nil)
