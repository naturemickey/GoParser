package ast

type Type_ struct {
	BaseNode
}

func (s *Type_) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Type_) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Type_) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Type_)(nil)
