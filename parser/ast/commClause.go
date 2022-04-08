package ast

type CommClause struct {
	BaseNode

	commCase      *CommCase
	statementList *StatementList
}

func NewCommClause(commCase *CommCase, statementList *StatementList) *CommClause {
	return &CommClause{commCase: commCase, statementList: statementList}
}

func (s *CommClause) CommCase() *CommCase {
	return s.commCase
}

func (s *CommClause) SetCommCase(commCase *CommCase) {
	s.commCase = commCase
}

func (s *CommClause) StatementList() *StatementList {
	return s.statementList
}

func (s *CommClause) SetStatementList(statementList *StatementList) {
	s.statementList = statementList
}

func (s *CommClause) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendNode(s.commCase).appendString(": ")
	if s.statementList != nil {
		cb.appendNode(s.statementList)
	}
	return cb
}

func (s *CommClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *CommClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*CommClause)(nil)
