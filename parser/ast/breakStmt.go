package ast

type BreakStmt struct {
	BaseNode
	label string
}

func (s *BreakStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *BreakStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("break")
	if s.label != "" {
		cb.blank().appendString(s.label)
	}
	return cb
}

func (s *BreakStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *BreakStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*BreakStmt)(nil)
