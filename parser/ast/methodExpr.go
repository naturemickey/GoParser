package ast

type MethodExpr struct {
	BaseNode
}

func (s *MethodExpr) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *MethodExpr) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *MethodExpr) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*MethodExpr)(nil)
