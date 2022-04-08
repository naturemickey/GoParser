package ast

type ConstDecl struct {
	BaseNode
	constSpecs []*ConstSpec
}

func (s *ConstDecl) ConstSpecs() []*ConstSpec {
	return s.constSpecs
}

func (s *ConstDecl) SetConstSpecs(constSpecs []*ConstSpec) {
	s.constSpecs = constSpecs
}

func NewConstDecl(constSpecs []*ConstSpec) *ConstDecl {
	return &ConstDecl{constSpecs: constSpecs}
}

func (s *ConstDecl) _Statement_() {
	//TODO implement me
	panic("implement me")
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
	if len(s.constSpecs) == 1 {
		cb.appendNode(s.constSpecs[0])
	} else {
		cb.appendString("(").newLine()
		for _, spec := range s.constSpecs {
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
