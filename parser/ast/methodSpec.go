package ast

type MethodSpec struct {
	BaseNode
}

func (s *MethodSpec) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *MethodSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *MethodSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*MethodSpec)(nil)
