package ast

type PrimaryExpr struct {
	BaseNode

	operand       *Operand
	conversion    *Conversion
	methodExpr    *MethodExpr
	primaryExpr   *PrimaryExpr
	id            string
	index         *Index
	slice         *Slice
	typeAssertion *TypeAssertion
	arguments     *Arguments
}

func (s *PrimaryExpr) Id() string {
	return s.id
}

func (s *PrimaryExpr) SetId(id string) {
	s.id = id
}

func (s *PrimaryExpr) Operand() *Operand {
	return s.operand
}

func (s *PrimaryExpr) SetOperand(operand *Operand) {
	s.operand = operand
}

func (s *PrimaryExpr) Conversion() *Conversion {
	return s.conversion
}

func (s *PrimaryExpr) SetConversion(conversion *Conversion) {
	s.conversion = conversion
}

func (s *PrimaryExpr) MethodExpr() *MethodExpr {
	return s.methodExpr
}

func (s *PrimaryExpr) SetMethodExpr(methodExpr *MethodExpr) {
	s.methodExpr = methodExpr
}

func (s *PrimaryExpr) PrimaryExpr() *PrimaryExpr {
	return s.primaryExpr
}

func (s *PrimaryExpr) SetPrimaryExpr(primaryExpr *PrimaryExpr) {
	s.primaryExpr = primaryExpr
}

func (s *PrimaryExpr) Index() *Index {
	return s.index
}

func (s *PrimaryExpr) SetIndex(index *Index) {
	s.index = index
}

func (s *PrimaryExpr) Slice() *Slice {
	return s.slice
}

func (s *PrimaryExpr) SetSlice(slice *Slice) {
	s.slice = slice
}

func (s *PrimaryExpr) TypeAssertion() *TypeAssertion {
	return s.typeAssertion
}

func (s *PrimaryExpr) SetTypeAssertion(typeAssertion *TypeAssertion) {
	s.typeAssertion = typeAssertion
}

func (s *PrimaryExpr) Arguments() *Arguments {
	return s.arguments
}

func (s *PrimaryExpr) SetArguments(arguments *Arguments) {
	s.arguments = arguments
}

func (s *PrimaryExpr) codeBuilder() *CodeBuilder {
	if s.operand != nil {
		return s.operand.codeBuilder()
	}
	if s.conversion != nil {
		return s.conversion.codeBuilder()
	}
	if s.methodExpr != nil {
		return s.methodExpr.codeBuilder()
	}
	cb := NewCodeBuilder()
	cb.appendNode(s.primaryExpr)
	if s.id != "" {
		cb.appendString(".").appendString(s.id)
	} else if s.index != nil {
		cb.appendNode(s.index)
	} else if s.slice != nil {
		cb.appendNode(s.slice)
	} else if s.typeAssertion != nil {
		cb.appendNode(s.typeAssertion)
	} else if s.arguments != nil {
		cb.appendNode(s.arguments)
	}
	return cb
}

func (s *PrimaryExpr) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *PrimaryExpr) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*PrimaryExpr)(nil)
