package ast

type ExprSwitchCase struct {
	BaseNode
}

func (s *ExprSwitchCase) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ExprSwitchCase) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ExprSwitchCase) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ExprSwitchCase)(nil)
