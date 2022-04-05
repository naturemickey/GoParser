package ast

type ImportSpec struct {
	BaseNode
	alias      string
	importPath *ImportPath
}

func (s *ImportSpec) Alias() string {
	return s.alias
}

func (s *ImportSpec) SetAlias(alias string) {
	s.alias = alias
}

func (s *ImportSpec) ImportPath() *ImportPath {
	return s.importPath
}

func (s *ImportSpec) SetImportPath(importPath *ImportPath) {
	s.importPath = importPath
}

func (s *ImportSpec) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.alias != "" {
		cb.appendString(s.alias).appendString(" ")
	}
	cb.appendNode(s.importPath)
	return cb
}

func (s *ImportSpec) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *ImportSpec) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*ImportSpec)(nil)
