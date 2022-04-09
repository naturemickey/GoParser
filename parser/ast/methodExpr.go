package ast

type MethodExpr struct {
	BaseNode
	nonNamedType *NonNamedType
	id           string
}

func NewMethodExpr(nonNamedType *NonNamedType, id string) *MethodExpr {
	return &MethodExpr{nonNamedType: nonNamedType, id: id}
}

func (s *MethodExpr) NonNamedType() *NonNamedType {
	return s.nonNamedType
}

func (s *MethodExpr) SetNonNamedType(nonNamedType *NonNamedType) {
	s.nonNamedType = nonNamedType
}

func (s *MethodExpr) Id() string {
	return s.id
}

func (s *MethodExpr) SetId(id string) {
	s.id = id
}

func (s *MethodExpr) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendNode(s.nonNamedType).appendString(".").appendString(s.id)
}

func (s *MethodExpr) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *MethodExpr) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*MethodExpr)(nil)
