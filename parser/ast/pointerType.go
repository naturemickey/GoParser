package ast

type PointerType struct {
	BaseNode

	type_ *Type_
}

func (s *PointerType) Type_() *Type_ {
	return s.type_
}

func (s *PointerType) SetType_(type_ *Type_) {
	s.type_ = type_
}

func (s *PointerType) _TypeLit_() {
	//TODO implement me
	panic("implement me")
}

func (s *PointerType) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("*").appendNode(s.type_)
}

func (s *PointerType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *PointerType) String() string {
	return s.codeBuilder().String()
}

var _ TypeLit = (*PointerType)(nil)
