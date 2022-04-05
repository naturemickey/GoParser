package ast

type MethodDecl struct {
	BaseNode
}

func (s *MethodDecl) _IFunctionMethodDeclaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *MethodDecl) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *MethodDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *MethodDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*MethodDecl)(nil)
var _ IFunctionMethodDeclaration = (*MethodDecl)(nil)
