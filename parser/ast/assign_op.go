package ast

type AssignOp struct {
	BaseNode
}

func (s *AssignOp) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *AssignOp) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *AssignOp) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*AssignOp)(nil)
