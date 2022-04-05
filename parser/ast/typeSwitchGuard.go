package ast

type TypeSwitchGuard struct {
	BaseNode

	id          string
	primaryExpr *PrimaryExpr
}

func (s *TypeSwitchGuard) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.id != "" {
		cb.appendString(s.id).appendString(" := ")
	}
	cb.appendNode(s.primaryExpr).appendString(".(type)")
	return cb
}

func (s *TypeSwitchGuard) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeSwitchGuard) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*TypeSwitchGuard)(nil)
