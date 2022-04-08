package ast

type ContinueStmt struct {
	BaseNode
	label string
}

func NewContinueStmt(label string) *ContinueStmt {
	return &ContinueStmt{label: label}
}

func (s *ContinueStmt) Label() string {
	return s.label
}

func (s *ContinueStmt) SetLabel(label string) {
	s.label = label
}

func (s *ContinueStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *ContinueStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("continue")
	if s.label != "" {
		cb.blank().appendString(s.label)
	}
	return cb
}

func (s *ContinueStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ContinueStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*ContinueStmt)(nil)
