package ast

type TypeSwitchGuard struct {
	BaseNode
}

func (s *TypeSwitchGuard) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchGuard) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchGuard) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeSwitchGuard)(nil)
