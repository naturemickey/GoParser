package ast

type Block struct {
	BaseNode
}

func (s *Block) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Block) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Block) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Block)(nil)
