package ast

type ExpressionStmt interface {
	SimpleStmt
	_ExpressionStmt_()
}
