package ast

type ChannelType struct {
	BaseNode
}

func (s *ChannelType) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *ChannelType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ChannelType) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ChannelType)(nil)
