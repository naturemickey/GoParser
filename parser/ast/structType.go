package ast

type StructType struct {
	BaseNode

	fieldDecls []*FieldDecl
}

func (s *StructType) FieldDecls() []*FieldDecl {
	return s.fieldDecls
}

func (s *StructType) SetFieldDecls(fieldDecls []*FieldDecl) {
	s.fieldDecls = fieldDecls
}

func (s *StructType) _TypeLit_() {
	//TODO implement me
	panic("implement me")
}

func (s *StructType) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("struct {\n")
	for _, decl := range s.fieldDecls {
		cb.appendNode(decl).newLine()
	}
	cb.appendString("}")
	return cb
}

func (s *StructType) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *StructType) String() string {
	return s.codeBuilder().String()
}

var _ TypeLit = (*StructType)(nil)
