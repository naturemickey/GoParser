package ast

type Conversion struct {
	BaseNode
}

func (s *Conversion) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Conversion) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Conversion) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Conversion)(nil)
