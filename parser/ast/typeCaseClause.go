package ast

type TypeCaseClause struct {
	BaseNode

	typeSwitchCase *TypeSwitchCase
	statementList  *StatementList
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
