package ast

type StructType struct {
	BaseNode
}

func (s *StructType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *StructType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *StructType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*StructType)(nil)
