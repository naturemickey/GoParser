package ast

type ConstDecl struct {
	BaseNode
	constSpec []*ConstSpec
}

func (s *ConstDecl) _Statement_() {
	//TODO implement me
	panic("implement me")
}

func (s *ConstDecl) ConstSpec() []*ConstSpec {
	return s.constSpec
}

func (s *ConstDecl) SetConstSpec(constSpec []*ConstSpec) {
	s.constSpec = constSpec
}

func (s *ConstDecl) _IFunctionMethodDeclaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *ConstDecl) _Declaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *ConstDecl) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("const ")
	if len(s.constSpec) == 1 {
		cb.appendNode(s.constSpec[0])
	} else {
		cb.appendString("(").newLine()
		for _, spec := range s.constSpec {
			cb.tab().appendNode(spec).newLine()
		}
		cb.appendString(")")
	}
	return cb
}

func (s *ConstDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ConstDecl) String() string {
	return s.codeBuilder().String()
}

var _ Declaration = (*ConstDecl)(nil)
