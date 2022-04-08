package ast

type ForClause struct {
	BaseNode
	initStmt   SimpleStmt
	expression *Expression
	postStmt   SimpleStmt
}

func NewForClause(initStmt SimpleStmt, expression *Expression, postStmt SimpleStmt) *ForClause {
	return &ForClause{initStmt: initStmt, expression: expression, postStmt: postStmt}
}

func (s *ForClause) InitStmt() SimpleStmt {
	return s.initStmt
}

func (s *ForClause) SetInitStmt(initStmt SimpleStmt) {
	s.initStmt = initStmt
}

func (s *ForClause) Expression() *Expression {
	return s.expression
}

func (s *ForClause) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *ForClause) PostStmt() SimpleStmt {
	return s.postStmt
}

func (s *ForClause) SetPostStmt(postStmt SimpleStmt) {
	s.postStmt = postStmt
}

func (s *ForClause) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendNode(s.initStmt).appendString("; ").appendNode(s.expression).appendString("; ").appendNode(s.postStmt)
}

func (s *ForClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ForClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ForClause)(nil)
