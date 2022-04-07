package ast

type ImportDecl struct {
	BaseNode
	importSpecs []*ImportSpec
}

func NewImportDecl(importSpecs []*ImportSpec) *ImportDecl {
	return &ImportDecl{importSpecs: importSpecs}
}

func (s *ImportDecl) ImportSpecs() []*ImportSpec {
	return s.importSpecs
}

func (s *ImportDecl) SetImportSpecs(importSpecs []*ImportSpec) {
	s.importSpecs = importSpecs
}

func (s *ImportDecl) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder().appendString("import ")
	if len(s.importSpecs) == 1 {
		cb.appendNode(s.importSpecs[0])
	} else {
		cb.appendString("(").newLine()
		for _, spec := range s.importSpecs {
			cb.tab().appendNode(spec).newLine()
		}
		cb.appendString(")")
	}
	return cb
}

func (s *ImportDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ImportDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ImportDecl)(nil)
