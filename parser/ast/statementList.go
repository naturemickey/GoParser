package ast

type StatementList struct {
	BaseNode
	statements []Statement
}

func (s *StatementList) Statements() []Statement {
	return s.statements
}

func (s *StatementList) SetStatements(statements []Statement) {
	s.statements = statements
}

func (s *StatementList) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	for _, statement := range s.statements {
		cb.appendNode(statement).newLine()
	}
	return cb
}

func (s *StatementList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *StatementList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*StatementList)(nil)
