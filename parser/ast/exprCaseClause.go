package ast

type ExprCaseClause struct {
	BaseNode
	exprSwitchCase *ExprSwitchCase
	statementList  *StatementList
}

func (s *ExprCaseClause) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendNode(s.exprSwitchCase).appendString(":\n")
	if s.statementList != nil {
		cb.appendNode(s.statementList)
	}
	return cb
}

func (s *ExprCaseClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ExprCaseClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ExprCaseClause)(nil)
