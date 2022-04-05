package ast

type ForClause struct {
	BaseNode
	initStmt   SimpleStmt
	expression *Expression
	postStmt   SimpleStmt
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
