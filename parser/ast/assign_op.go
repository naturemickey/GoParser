package ast

type AssignOp struct {
	BaseNode
	assignPrefix string
}

func NewAssignOp(assignPrefix string) *AssignOp {
	return &AssignOp{assignPrefix: assignPrefix}
}

func (s *AssignOp) AssignPrefix() string {
	return s.assignPrefix
}

func (s *AssignOp) SetAssignPrefix(assignPrefix string) {
	s.assignPrefix = assignPrefix
}

func (s *AssignOp) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString(s.assignPrefix).appendString("=")
}

func (s *AssignOp) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *AssignOp) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*AssignOp)(nil)
