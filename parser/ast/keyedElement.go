package ast

type KeyedElement struct {
	BaseNode

	key     Key
	element Element
}

func NewKeyedElement(key Key, element Element) *KeyedElement {
	return &KeyedElement{key: key, element: element}
}

func (s *KeyedElement) Key() Key {
	return s.key
}

func (s *KeyedElement) SetKey(key Key) {
	s.key = key
}

func (s *KeyedElement) Element() Element {
	return s.element
}

func (s *KeyedElement) SetElement(element Element) {
	s.element = element
}

func (s *KeyedElement) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.key != nil {
		cb.appendNode(s.key).appendString(": ")
	}
	cb.appendNode(s.element)
	return cb
}

func (s *KeyedElement) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *KeyedElement) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*KeyedElement)(nil)
