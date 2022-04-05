package ast

type IncDecStmt struct {
	BaseNode
	expression *Expression
	isPlus     bool
}

func (s *IncDecStmt) IsPlus() bool {
	return s.isPlus
}

func (s *IncDecStmt) SetIsPlus(isPlus bool) {
	s.isPlus = isPlus
}

func (s *IncDecStmt) Expression() *Expression {
	return s.expression
}

func (s *IncDecStmt) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *IncDecStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *IncDecStmt) _SimpleStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *IncDecStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendNode(s.expression)
	if s.isPlus {
		cb.appendString("++")
	} else {
		cb.appendString("--")
	}
	return cb
}

func (s *IncDecStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *IncDecStmt) String() string {
	return s.codeBuilder().String()
}

var _ SimpleStmt = (*IncDecStmt)(nil)
