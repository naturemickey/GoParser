package ast

type PackageClause struct {
	BaseNode
}

func (s *PackageClause) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *PackageClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *PackageClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*PackageClause)(nil)
