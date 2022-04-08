package ast

type TypeList struct {
	BaseNode

	types []*Type_
}

func NewTypeList(types []*Type_) *TypeList {
	return &TypeList{types: types}
}

func (s *TypeList) Types() []*Type_ {
	return s.types
}

func (s *TypeList) SetTypes(types []*Type_) {
	s.types = types
}

func (s *TypeList) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	for _, type_ := range s.types {
		if type_ == nil {
			cb.appendString("nil")
		} else {
			cb.appendNode(type_)
		}
		cb.appendString(", ")
	}
	cb.deleteLast()
	return cb
}

func (s *TypeList) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeList) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeList)(nil)
