package ast

type LabeledStmt struct {
	BaseNode
}

func (s *LabeledStmt) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *LabeledStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *LabeledStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*LabeledStmt)(nil)
