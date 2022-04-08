package ast

type ExprCaseClause struct {
	BaseNode
	exprSwitchCase *ExprSwitchCase
	statementList  *StatementList
}

func NewExprCaseClause(exprSwitchCase *ExprSwitchCase, statementList *StatementList) *ExprCaseClause {
	return &ExprCaseClause{exprSwitchCase: exprSwitchCase, statementList: statementList}
}

func (s *ExprCaseClause) ExprSwitchCase() *ExprSwitchCase {
	return s.exprSwitchCase
}

func (s *ExprCaseClause) SetExprSwitchCase(exprSwitchCase *ExprSwitchCase) {
	s.exprSwitchCase = exprSwitchCase
}

func (s *ExprCaseClause) StatementList() *StatementList {
	return s.statementList
}

func (s *ExprCaseClause) SetStatementList(statementList *StatementList) {
	s.statementList = statementList
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
