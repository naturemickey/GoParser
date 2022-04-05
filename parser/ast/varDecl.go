package ast

type VarDecl struct {
	BaseNode
	varSpecs []*VarSpec
}

func (s *VarDecl) VarSpecs() []*VarSpec {
	return s.varSpecs
}

func (s *VarDecl) SetVarSpecs(varSpecs []*VarSpec) {
	s.varSpecs = varSpecs
}

func (s *VarDecl) _IFunctionMethodDeclaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *VarDecl) _Declaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *VarDecl) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	cb.appendString("var ")
	if len(s.varSpecs) == 1 {
		cb.appendNode(s.varSpecs[0])
	} else {
		cb.appendString("(")
		for _, spec := range s.varSpecs {
			cb.appendNode(spec).newLine()
		}
		cb.appendString(")")
	}
	return cb
}

func (s *VarDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *VarDecl) String() string {
	return s.codeBuilder().String()
}

var _ Declaration = (*VarDecl)(nil)
