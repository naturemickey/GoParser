package ast

type LiteralValue struct {
	BaseNode
	elementList *ElementList
}

func (s *LiteralValue) _Element_() {
	//TODO implement me
	panic("implement me")
}

func (s *LiteralValue) _Key_() {
	//TODO implement me
	panic("implement me")
}

func (s *LiteralValue) ElementList() *ElementList {
	return s.elementList
}

func (s *LiteralValue) SetElementList(elementList *ElementList) {
	s.elementList = elementList
}

func (s *LiteralValue) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("{").appendNode(s.elementList).appendString("}")
}

func (s *LiteralValue) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *LiteralValue) String() string {
	return s.codeBuilder().String()
}

var _ Key = (*LiteralValue)(nil)
var _ Element = (*LiteralValue)(nil)
