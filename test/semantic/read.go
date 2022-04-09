package semantic

import "GoParser/parser/ast"

func Read (tree ast.SourceFile)  {
	scope := NewScopeRoot()
	for _, fmd := range tree.FunctionOrMethodOrDeclarations() {
		switch child := fmd.(type) {
		case *ast.FunctionDecl:
			name        :=child.FunName()
			tree      :=child
			type_       := FUNCATION
			typeLiteral string

		case *ast.MethodDecl:
		case ast.Declaration:
		}
	}
}