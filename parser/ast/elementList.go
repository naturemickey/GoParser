package ast

type ElementList struct {
	BaseNode

	keyedElement []*KeyedElement
}

func (s *ElementList) KeyedElement() []*KeyedElement {
	return s.keyedElement
}

func (s *ElementList) SetKeyedElement(keyedElement []*KeyedElement) {
	s.keyedElement = keyedElement
}

func (s *ElementList) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	for _, element := range s.keyedElement {
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
