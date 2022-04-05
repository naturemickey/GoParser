package ast

type ForStmt struct {
	BaseNode

	expression  *Expression
	forClause   *ForClause
	rangeClause *RangeClause
	block       *Block
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
