package ast

type LabeledStmt struct {
	BaseNode
	label     string
	statement Statement
}

func (s *LabeledStmt) Label() string {
	return s.label
}

func (s *LabeledStmt) SetLabel(label string) {
	s.label = label
}

func (s *LabeledStmt) Statement() Statement {
	return s.statement
}

func (s *LabeledStmt) SetStatement(statement Statement) {
	s.statement = statement
}

func (s *LabeledStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *LabeledStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()

	cb.appendString(s.label).appendString(":")
	if s.statement != nil {
		cb.appendString(" ").appendNode(s.statement)
	}
	return cb
}

func (s *LabeledStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *LabeledStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*LabeledStmt)(nil)
