package ast

type Result struct {
	BaseNode
}

func (s *Result) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Result) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Result) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Result)(nil)
