package ast

type SimpleStmt interface {
	Statement
	_SimpleStmt_()
}
