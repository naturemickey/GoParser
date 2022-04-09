package ast

type InterfaceType struct {
	BaseNode

	methodSpecOrTypeNames []IMethodSpecOrTypeName
}

func NewInterfaceType(methodSpecOrTypeNames []IMethodSpecOrTypeName) *InterfaceType {
	return &InterfaceType{methodSpecOrTypeNames: methodSpecOrTypeNames}
}

func (s *InterfaceType) MethodSpecOrTypeNames() []IMethodSpecOrTypeName {
	return s.methodSpecOrTypeNames
}

func (s *InterfaceType) SetMethodSpecOrTypeNames(methodSpecOrTypeNames []IMethodSpecOrTypeName) {
	s.methodSpecOrTypeNames = methodSpecOrTypeNames
}

func (s *InterfaceType) _TypeLit_() {
	//TODO implement me
	panic("implement me")
}

func (s *InterfaceType) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("interface {\n")
	for _, mt := range s.methodSpecOrTypeNames {
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
