package ast

type ArrayType struct {
	BaseNode
}

func (s *ArrayType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ArrayType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ArrayType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ArrayType)(nil)
