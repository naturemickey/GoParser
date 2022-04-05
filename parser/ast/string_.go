package ast

type String_ struct {
	BaseNode
}

func (s *String_) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *String_) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *String_) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*String_)(nil)
