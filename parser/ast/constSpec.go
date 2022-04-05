package ast

type ConstSpec struct {
	BaseNode
	identifierList *IdentifierList
	type_          *Type_
	expressionList *ExpressionList
}

func (s *ConstSpec) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendNode(s.identifierList).blank()

	if s.expressionList != nil {
		if s.type_ != nil {
			cb.appendNode(s.type_).blank()
		}
		cb.appendString("= ")
		cb.appendNode(s.expressionList)
	}
	return cb
}

func (s *ConstSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ConstSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ConstSpec)(nil)
