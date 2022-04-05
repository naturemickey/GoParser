package ast

type PackageClause struct {
}

func (s PackageClause) String() {
	//TODO implement me
	panic("implement me")
}

func (s PackageClause) Children() []INode {
	//TODO implement me
	panic("implement me")
}

var _ INode = (*PackageClause)(nil)
