package ast

type VarSpec struct {
	BaseNode
}

func (s *VarSpec) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *VarSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *VarSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*VarSpec)(nil)
