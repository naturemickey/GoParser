package ast

type RecvStmt struct {
	BaseNode

	expressionList *ExpressionList
	identifierList *IdentifierList
	recvExpr       *Expression
}

func NewRecvStmt(expressionList *ExpressionList, identifierList *IdentifierList, recvExpr *Expression) *RecvStmt {
	return &RecvStmt{expressionList: expressionList, identifierList: identifierList, recvExpr: recvExpr}
}

func (s *RecvStmt) ExpressionList() *ExpressionList {
	return s.expressionList
}

func (s *RecvStmt) SetExpressionList(expressionList *ExpressionList) {
	s.expressionList = expressionList
}

func (s *RecvStmt) IdentifierList() *IdentifierList {
	return s.identifierList
}

func (s *RecvStmt) SetIdentifierList(identifierList *IdentifierList) {
	s.identifierList = identifierList
}

func (s *RecvStmt) RecvExpr() *Expression {
	return s.recvExpr
}

func (s *RecvStmt) SetRecvExpr(recvExpr *Expression) {
	s.recvExpr = recvExpr
}

func (s *RecvStmt) codeBuilder() *CodeBuilder {
	cb := NewCodeBuilder()
	if s.expressionList != nil {
		cb.appendNode(s.expressionList).appendString(" = ")
	} else if s.identifierList != nil {
		cb.appendNode(s.identifierList).appendString(" := ")
	}
	cb.appendNode(s.recvExpr)
	return cb
}

func (s *RecvStmt) Children() []INode {
	//TODO implement me
	panic("implement me")
}

func (s *RecvStmt) String() string {
	return s.codeBuilder().String()
}

var _ INode = (*RecvStmt)(nil)
