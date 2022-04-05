package ast

type CommCase struct {
	BaseNode
}

func (s *CommCase) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *CommCase) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *CommCase) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*CommCase)(nil)
