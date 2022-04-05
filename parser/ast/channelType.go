package ast

type ChannelType struct {
}

func (s ChannelType) String() {
	//TODO implement me
	panic("implement me")
}

func (s ChannelType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*ChannelType)(nil)
