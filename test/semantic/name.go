package semantic

import "GoParser/parser/ast"

type Name struct {
	name        string
	tree        ast.INode
	type_       NameType
	typeLiteral string
}

func (n Name) Name() string {
	return n.name
}

func (n Name) Tree() ast.INode {
	return n.tree
}

func (n Name) Type_() NameType {
	return n.type_
}

func (n Name) TypeLiteral() string {
	return n.typeLiteral
}

func NewName(name string, tree ast.INode, type_ NameType, typeLiteral string) *Name {
	return &Name{name: name, tree: tree, type_: type_, typeLiteral: typeLiteral}
}

type NameType int

const (
	FUNCATION = iota
	METHOD
	VARIABLE
)
