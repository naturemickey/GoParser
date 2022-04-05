package ast

type PackageClause struct {
	BaseNode
	packageName string
}

func NewPackageClause(packageName string) *PackageClause {
	return &PackageClause{packageName: packageName}
}

func (s *PackageClause) PackageName() string {
	return s.packageName
}

func (s *PackageClause) SetPackageName(packageName string) {
	s.packageName = packageName
}

func (s *PackageClause) codeBuilder() *CodeBuilder {
	return NewCodeBuilder().appendString("package ").appendString(s.packageName)
}

func (s *PackageClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *PackageClause) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*PackageClause)(nil)
