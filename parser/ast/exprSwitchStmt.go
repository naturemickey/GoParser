package ast

type ExprSwitchStmt struct {
	BaseNode
	expression      *Expression
	simpleStmt      SimpleStmt
	exprCaseClauses []*ExprCaseClause
}

func (s *ExprSwitchStmt) ExprCaseClauses() []*ExprCaseClause {
	return s.exprCaseClauses
}

func (s *ExprSwitchStmt) SetExprCaseClauses(exprCaseClauses []*ExprCaseClause) {
	s.exprCaseClauses = exprCaseClauses
}

func (s *ExprSwitchStmt) Expression() *Expression {
	return s.expression
}

func (s *ExprSwitchStmt) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *ExprSwitchStmt) SimpleStmt() SimpleStmt {
	return s.simpleStmt
}

func (s *ExprSwitchStmt) SetSimpleStmt(simpleStmt SimpleStmt) {
	s.simpleStmt = simpleStmt
}

func (s *ExprSwitchStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *ExprSwitchStmt) _SwitchStmt_() {
	//TODO implement me
	panic("implement me")
}

func (s *ExprSwitchStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("switch ")
	if s.simpleStmt != nil {
		cb.appendNode(s.simpleStmt).appendString("; ")
	}
	cb.appendNode(s.expression)
	cb.appendString("{").newLine()
	for _, clause := range s.exprCaseClauses {
		cb.appendNode(clause).newLine()
	}
	cb.appendString("}")
	return cb
}

func (s *ExprSwitchStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ExprSwitchStmt) String() string {
	return s.codeBuilder().String()
}

var _ SwitchStmt = (*ExprSwitchStmt)(nil)
