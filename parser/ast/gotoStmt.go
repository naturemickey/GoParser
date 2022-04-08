package ast

type GotoStmt struct {
	BaseNode
	label string
}

func NewGotoStmt(label string) *GotoStmt {
	return &GotoStmt{label: label}
}

func (s *GotoStmt) Label() string {
	return s.label
}

func (s *GotoStmt) SetLabel(label string) {
	s.label = label
}

func (s *GotoStmt) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *GotoStmt) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("goto ").appendString(s.label)
}

func (s *GotoStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *GotoStmt) String() string {
	return s.codeBuilder().String()
}

var _ Statement = (*GotoStmt)(nil)
