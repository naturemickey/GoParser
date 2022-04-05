package ast

type ImportDecl struct {
	BaseNode
	importSpec []*ImportSpec
}

func (s *ImportDecl) ImportSpec() []*ImportSpec {
	return s.importSpec
}

func (s *ImportDecl) SetImportSpec(importSpec []*ImportSpec) {
	s.importSpec = importSpec
}

func (s *ImportDecl) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder().appendString("import ")
	if len(s.importSpec) == 1 {
		cb.appendNode(s.importSpec[0])
	} else {
		cb.appendString("(").newLine()
		for _, spec := range s.importSpec {
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
