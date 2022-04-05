package ast

type INode interface {
	String() string
	Children() []INode
	codeBuilder() *CodeBuilder
}
