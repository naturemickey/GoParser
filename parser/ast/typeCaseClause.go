package ast

type TypeCaseClause struct {
	BaseNode

	typeSwitchCase *TypeSwitchCase
	statementList  *StatementList
}

func NewTypeCaseClause(typeSwitchCase *TypeSwitchCase, statementList *StatementList) *TypeCaseClause {
	return &TypeCaseClause{typeSwitchCase: typeSwitchCase, statementList: statementList}
}

func (s *TypeCaseClause) TypeSwitchCase() *TypeSwitchCase {
	return s.typeSwitchCase
}

func (s *TypeCaseClause) SetTypeSwitchCase(typeSwitchCase *TypeSwitchCase) {
	s.typeSwitchCase = typeSwitchCase
}

func (s *TypeCaseClause) StatementList() *StatementList {
	return s.statementList
}

func (s *TypeCaseClause) SetStatementList(statementList *StatementList) {
	s.statementList = statementList
}

func (s *TypeCaseClause) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendNode(s.typeSwitchCase).appendString(": ")
	if s.statementList != nil {
		cb.newLine().appendNode(s.statementList)
	}
	return cb
}

func (s *TypeCaseClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeCaseClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeCaseClause)(nil)
