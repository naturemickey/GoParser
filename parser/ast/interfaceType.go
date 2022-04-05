package ast

type InterfaceType struct {
	BaseNode

	methodSpecOrTypeName []IMethodSpecOrTypeName
}

func (s *InterfaceType) MethodSpecOrTypeName() []IMethodSpecOrTypeName {
	return s.methodSpecOrTypeName
}

func (s *InterfaceType) SetMethodSpecOrTypeName(methodSpecOrTypeName []IMethodSpecOrTypeName) {
	s.methodSpecOrTypeName = methodSpecOrTypeName
}

func (s *InterfaceType) _TypeLit_() {
	//TODO implement me
	panic("implement me")
}

func (s *InterfaceType) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("interface {\n")
	for _, mt := range s.methodSpecOrTypeName {
		cb.appendNode(mt).newLine()
	}
	cb.appendString("}")
	return cb
}

func (s *InterfaceType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *InterfaceType) String() string {
	return s.codeBuilder().String()
}

var _ TypeLit = (*InterfaceType)(nil)
