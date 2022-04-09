package ast

type ChannelType struct {
	BaseNode

	title       ChannelTitle
	elementType ElementType
}

func NewChannelType(title ChannelTitle, elementType ElementType) *ChannelType {
	return &ChannelType{title: title, elementType: elementType}
}

type ChannelTitle int

const (
	CHAN         = 0
	CHAN_RECEIVE = 1
	RECEIVE_CHAN = 2
)

func (s *ChannelType) Title() ChannelTitle {
	return s.title
}

func (s *ChannelType) SetTitle(title ChannelTitle) {
	s.title = title
}

func (s *ChannelType) ElementType() ElementType {
	return s.elementType
}

func (s *ChannelType) SetElementType(elementType ElementType) {
	s.elementType = elementType
}

func (s *ChannelType) _TypeLit_() {
	//TODO implement me
	panic("implement me")
}

func (s *ChannelType) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	switch s.title {
	case CHAN:
		cb.appendString("chan ")
	case CHAN_RECEIVE:
		cb.appendString("chan <- ")
	case RECEIVE_CHAN:
		cb.appendString("<- chan ")
	}
	cb.appendNode(s.elementType)
	return cb
}

func (s *ChannelType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ChannelType) String() string {
	return s.codeBuilder().String()
}

var _ TypeLit = (*ChannelType)(nil)
