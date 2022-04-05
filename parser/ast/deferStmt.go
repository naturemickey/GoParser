package ast

type DeferStmt struct {
	BaseNode
	expression *Expression
}

func (s *DeferStmt) Expression() *Expression {
	return s.expression
}

func (s *DeferStmt) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *DeferStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *DeferStmt) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("defer ").appendNode(s.expression)
}

func (s *DeferStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *DeferStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*DeferStmt)(nil)
