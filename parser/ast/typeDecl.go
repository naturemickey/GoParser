package ast

type TypeDecl struct {
	BaseNode
	typeSpecs []*TypeSpec
}

func (s *TypeDecl) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *TypeDecl) _IFunctionMethodDeclaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *TypeDecl) _Declaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *TypeDecl) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("type ")
	if len(s.typeSpecs) == 1 {
		cb.appendNode(s.typeSpecs[0])
	} else {
		for _, spec := range s.typeSpecs {
			cb.tab().appendNode(spec).newLine()
		}
	}
	return cb
}

func (s *TypeDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *TypeDecl) String() string {
	return s.codeBuilder().String()
}

var _ Declaration = (*TypeDecl)(nil)
