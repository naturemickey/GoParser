package ast

type SourceFile struct {
	BaseNode
	packageClause                 *PackageClause
	importDecl                    []*ImportDecl
	functionOrMethodOrDeclaration []IFunctionMethodDeclaration
	//functionDecl  *FunctionDecl
	//methodDecl    *MethodDecl
	//declaration   *Declaration
}

func (s *SourceFile) FunctionOrMethodOrDeclaration() []IFunctionMethodDeclaration {
	return s.functionOrMethodOrDeclaration
}

func (s *SourceFile) SetFunctionOrMethodOrDeclaration(functionOrMethodOrDeclaration []IFunctionMethodDeclaration) {
	s.functionOrMethodOrDeclaration = functionOrMethodOrDeclaration
}

func (s *SourceFile) PackageClause() *PackageClause {
	return s.packageClause
}

func (s *SourceFile) SetPackageClause(packageClause *PackageClause) {
	s.packageClause = packageClause
}

func (s *SourceFile) ImportDecl() []*ImportDecl {
	return s.importDecl
}

func (s *SourceFile) SetImportDecl(importDecl []*ImportDecl) {
	s.importDecl = importDecl
}

func (s *SourceFile) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SourceFile) codeBuilder() *CodeBuilder {
	cb := new(CodeBuilder)

	cb.appendNode(s.packageClause)

	for _, decl := range s.importDecl {
		cb.appendNode(decl).newLine()
	}

	for _, fmd := range s.functionOrMethodOrDeclaration {
		cb.appendNode(fmd).newLine()
	}

	return cb
}

func (s *SourceFile) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*SourceFile)(nil)
