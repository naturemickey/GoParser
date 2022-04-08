package ast

type TypeSwitchGuard struct {
	BaseNode

	id          string
	primaryExpr *PrimaryExpr
}

func NewTypeSwitchGuard(id string, primaryExpr *PrimaryExpr) *TypeSwitchGuard {
	return &TypeSwitchGuard{id: id, primaryExpr: primaryExpr}
}

func (s *TypeSwitchGuard) Id() string {
	return s.id
}

func (s *TypeSwitchGuard) SetId(id string) {
	s.id = id
}

func (s *TypeSwitchGuard) PrimaryExpr() *PrimaryExpr {
	return s.primaryExpr
}

func (s *TypeSwitchGuard) SetPrimaryExpr(primaryExpr *PrimaryExpr) {
	s.primaryExpr = primaryExpr
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
