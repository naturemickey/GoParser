package ast

type SourceFile struct {
	BaseNode
	packageClause                  *PackageClause
	importDecls                    []*ImportDecl
	functionOrMethodOrDeclarations []IFunctionMethodDeclaration
	//functionDecl  *FunctionDecl
	//methodDecl    *MethodDecl
	//declaration   *Declaration
}

func NewSourceFile(packageClause *PackageClause, importDecls []*ImportDecl, functionOrMethodOrDeclarations []IFunctionMethodDeclaration) *SourceFile {
	return &SourceFile{packageClause: packageClause, importDecls: importDecls, functionOrMethodOrDeclarations: functionOrMethodOrDeclarations}
}

func (s *SourceFile) addFunctionDecl(d *FunctionDecl) {
	s.functionOrMethodOrDeclarations = append(s.functionOrMethodOrDeclarations, d)
}
func (s *SourceFile) addMethodDecl(d *MethodDecl) {
	s.functionOrMethodOrDeclarations = append(s.functionOrMethodOrDeclarations, d)
}
func (s *SourceFile) addDeclaration(d Declaration) {
	s.functionOrMethodOrDeclarations = append(s.functionOrMethodOrDeclarations, d)
}

func (s *SourceFile) ImportDecls() []*ImportDecl {
	return s.importDecls
}

func (s *SourceFile) SetImportDecls(importDecls []*ImportDecl) {
	s.importDecls = importDecls
}

func (s *SourceFile) FunctionOrMethodOrDeclarations() []IFunctionMethodDeclaration {
	return s.functionOrMethodOrDeclarations
}

func (s *SourceFile) SetFunctionOrMethodOrDeclarations(functionOrMethodOrDeclarations []IFunctionMethodDeclaration) {
	s.functionOrMethodOrDeclarations = functionOrMethodOrDeclarations
}

func (s *SourceFile) PackageClause() *PackageClause {
	return s.packageClause
}

func (s *SourceFile) SetPackageClause(packageClause *PackageClause) {
	s.packageClause = packageClause
}

func (s *SourceFile) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *SourceFile) codeBuilder() *CodeBuilder {
	cb := new(CodeBuilder)

	cb.appendNode(s.packageClause).newLine()

	for _, decl := range s.importDecls {
		cb.appendNode(decl).newLine()
	}

	for _, fmd := range s.functionOrMethodOrDeclarations {
		cb.appendNode(fmd).newLine()
	}

	return cb
}

func (s *SourceFile) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*SourceFile)(nil)
