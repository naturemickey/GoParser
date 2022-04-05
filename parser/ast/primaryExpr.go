package ast

type PrimaryExpr struct {
	BaseNode
}

func (s *PrimaryExpr) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *PrimaryExpr) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *PrimaryExpr) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*PrimaryExpr)(nil)
