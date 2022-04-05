package ast

type Declaration interface {
	IFunctionMethodDeclaration
	Statement
	_Declaration_()
}
