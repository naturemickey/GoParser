package ast

type OperandName struct {
	BaseNode
}

func (s *OperandName) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *OperandName) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *OperandName) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*OperandName)(nil)
