package ast

type Statement interface {
	INode
	_Statement_()
}
