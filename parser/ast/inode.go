package ast

type INode interface {
	String()
	Children() []INode
}
