package ast

type FunctionDecl struct {
	BaseNode
}

func (s *FunctionDecl) _IFunctionMethodDeclaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionDecl) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionDecl) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *FunctionDecl) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*FunctionDecl)(nil)
var _ IFunctionMethodDeclaration = (*FunctionDecl)(nil)
