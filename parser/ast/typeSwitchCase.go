package ast

type TypeSwitchCase struct {
	BaseNode

	typeList *TypeList
}

func NewTypeSwitchCase(typeList *TypeList) *TypeSwitchCase {
	return &TypeSwitchCase{typeList: typeList}
}

func (s *TypeSwitchCase) TypeList() *TypeList {
	return s.typeList
}

func (s *TypeSwitchCase) SetTypeList(typeList *TypeList) {
	s.typeList = typeList
}

func (s *TypeSwitchCase) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.typeList != nil {
		cb.appendString("case ").appendNode(s.typeList)
	} else {
		cb.appendString("default")
	}
	return cb
}

func (s *TypeSwitchCase) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchCase) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeSwitchCase)(nil)
