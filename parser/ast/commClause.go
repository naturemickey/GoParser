package ast

type CommClause struct {
	BaseNode

	commCase      *CommCase
	statementList *StatementList
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
