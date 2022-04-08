package ast

type GoStmt struct {
	BaseNode

	expression *Expression
}

func NewGoStmt(expression *Expression) *GoStmt {
	return &GoStmt{expression: expression}
}

func (s *GoStmt) Expression() *Expression {
	return s.expression
}

func (s *GoStmt) SetExpression(expression *Expression) {
	s.expression = expression
}

func (s *GoStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *GoStmt) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("go ").appendNode(s.expression)
}

func (s *GoStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *GoStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*GoStmt)(nil)
