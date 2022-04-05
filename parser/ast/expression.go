package ast

type Expression struct {
	BaseNode
}

func (s *Expression) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Expression) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Expression)(nil)
