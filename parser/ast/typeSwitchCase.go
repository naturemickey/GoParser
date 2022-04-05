package ast

type TypeSwitchCase struct {
	BaseNode
}

func (s *TypeSwitchCase) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchCase) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchCase) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeSwitchCase)(nil)
