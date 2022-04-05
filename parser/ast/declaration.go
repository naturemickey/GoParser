package ast

type Declaration struct {
	BaseNode
}

func (s *Declaration) _IFunctionMethodDeclaration_() {
	//TODO implement me
	panic("implement me")
}

func (s *Declaration) codeBuilder() *CodeBuilder {
	//TODO implement me
	panic("implement me")
}

func (s *Declaration) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *Declaration) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*Declaration)(nil)
var _ IFunctionMethodDeclaration = (*Declaration)(nil)
