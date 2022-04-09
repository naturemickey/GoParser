package ast

type ElementList struct {
	BaseNode

	keyedElements []*KeyedElement
}

func NewElementList(keyedElements []*KeyedElement) *ElementList {
	return &ElementList{keyedElements: keyedElements}
}

func (s *ElementList) KeyedElements() []*KeyedElement {
	return s.keyedElements
}

func (s *ElementList) SetKeyedElements(keyedElements []*KeyedElement) {
	s.keyedElements = keyedElements
}

func (s *ElementList) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	for _, element := range s.keyedElements {
		cb.appendNode(element).appendString(", ")
	}
	cb.deleteLast()
	return cb
}

func (s *ElementList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ElementList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ElementList)(nil)
